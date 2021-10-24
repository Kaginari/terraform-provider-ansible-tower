package tower

import (
	"context"
	"fmt"
	tower "github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		UpdateContext: resourceProjectUpdate,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of this project",
			},

			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Optional description of this project.",
			},

			"local_path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
			},

			"scm_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "One of \"\" (manual), git, hg, svn",
			},

			"scm_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},
			"scm_credential_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Numeric ID of the scm used credential",
			},
			"scm_branch": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specific branch, tag or commit to checkout.",
			},
			"scm_clean": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scm_delete_on_update": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"organisation_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Numeric ID of the project organization",
			},
			"scm_update_on_launch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scm_update_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.ProjectService

	orgID := d.Get("organisation_id").(int)
	projectName := d.Get("name").(string)
	_, res, err := awxService.ListProjects(map[string]string{
		"name":         projectName,
		"organization": strconv.Itoa(orgID),
	},
	)
	if err != nil {
		return DiagnosticsMessage("Create: Fail to find Project", fmt.Sprintf("Fail to find Project %s Organisation ID %v, %s", projectName, orgID, err.Error()))
	}
	if len(res.Results) >= 1 {
		return DiagnosticsMessage("Create: Allways exist", fmt.Sprintf("Project with name %s  already exists in the Organisation ID %v", projectName, orgID))
	}
	credentials := ""
	if d.Get("scm_credential_id").(int) > 0 {
		credentials = strconv.Itoa(d.Get("scm_credential_id").(int))
	}
	result, err := awxService.CreateProject(map[string]interface{}{
		"name":                 projectName,
		"description":          d.Get("description").(string),
		"local_path":           d.Get("local_path").(string),
		"scm_type":             d.Get("scm_type").(string),
		"scm_url":              d.Get("scm_url").(string),
		"scm_branch":           d.Get("scm_branch").(string),
		"scm_clean":            d.Get("scm_clean").(bool),
		"scm_delete_on_update": d.Get("scm_delete_on_update").(bool),
		"organization":         d.Get("organisation_id").(int),
		"credential":           credentials,

		"scm_update_on_launch":     d.Get("scm_update_on_launch").(bool),
		"scm_update_cache_timeout": d.Get("scm_update_cache_timeout").(int),
	}, map[string]string{})
	if err != nil {
		return DiagnosticsMessage("Create: Project not created", fmt.Sprintf("Project with name %s  in the Organisation ID %v not created, %s", projectName, orgID, err.Error()))
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.ProjectService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}
	credentials := ""
	if d.Get("scm_credential_id").(int) > 0 {
		credentials = strconv.Itoa(d.Get("scm_credential_id").(int))
	}
	_, err = awxService.UpdateProject(id, map[string]interface{}{
		"name":                     d.Get("name").(string),
		"description":              d.Get("description").(string),
		"local_path":               d.Get("local_path").(string),
		"scm_type":                 d.Get("scm_type").(string),
		"scm_url":                  d.Get("scm_url").(string),
		"scm_branch":               d.Get("scm_branch").(string),
		"scm_clean":                d.Get("scm_clean").(bool),
		"scm_delete_on_update":     d.Get("scm_delete_on_update").(bool),
		"credential":               credentials,
		"organization":             d.Get("organisation_id").(int),
		"scm_update_on_launch":     d.Get("scm_update_on_launch").(bool),
		"scm_update_cache_timeout": d.Get("scm_update_cache_timeout").(int),
	}, map[string]string{})
	if err != nil {
		return DiagnosticsMessage("Update: Fail To Update Project",fmt.Sprintf("Fail to get Project with ID %v, got %s", id, err.Error()))
	}
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*tower.AWX)
	awxService := client.ProjectService

	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}

	res, err := awxService.GetProjectById(id, make(map[string]string))
	if err != nil {
		return DiagNotFoundFail("project", id, err)
	}
	d = setProjectResourceData(d, res)
	return diags
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	digMessagePart := "Project"
	client := m.(*tower.AWX)
	awxService := client.ProjectService
	var jobID int
	var finished time.Time
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}

	res, err := awxService.GetProjectById(id, make(map[string]string))
	if err != nil {
		d.SetId("")
		return DiagNotFoundFail("project", id, err)
	}

	if res.SummaryFields.CurrentJob["id"] != nil {
		jobID = int(res.SummaryFields.CurrentJob["id"].(float64))
	} else if res.SummaryFields.LastJob["id"] != nil {
		jobID = int(res.SummaryFields.LastJob["id"].(float64))
	}
	if jobID != 0 {
		_, err = client.ProjectUpdatesService.ProjectUpdateCancel(jobID)
		if err != nil {
			return DiagnosticsMessage(
				"Delete: Fail to canel Job",
				fmt.Sprintf("Fail to canel the Job %v for Project with ID %v, got %s",
				jobID, id, err.Error()),
			)
		}
	}
	// check if finished is 0
	for finished.IsZero() {
		prj, _ := client.ProjectUpdatesService.ProjectUpdateGet(jobID)
		finished = prj.Finished
		time.Sleep(1 * time.Second)
	}

	if _, err = awxService.DeleteProject(id); err != nil {
		return DiagDeleteFail(digMessagePart, fmt.Sprintf("ProjectID %v, got %s ", id, err.Error()))
	}
	d.SetId("")
	return diags
}

func setProjectResourceData(d *schema.ResourceData, r *tower.Project) *schema.ResourceData {
	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("scm_type", r.ScmType)
	d.Set("scm_url", r.ScmURL)
	d.Set("scm_branch", r.ScmBranch)
	d.Set("scm_clean", r.ScmClean)
	d.Set("scm_delete_on_update", r.ScmDeleteOnUpdate)
	d.Set("organisation_id", r.Organization)

	id, err := strconv.Atoi(r.Credential)
	if err == nil {
		d.Set("scm_credential_id", id)
	}
	d.Set("scm_update_on_launch", r.ScmUpdateOnLaunch)
	d.Set("scm_update_cache_timeout", r.ScmUpdateCacheTimeout)

	d.SetId(strconv.Itoa(r.ID))
	return d
}