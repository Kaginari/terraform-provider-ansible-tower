package tower

import (
	"context"
	"fmt"
	"github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInventorySource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventorySourceCreate,
		ReadContext:   resourceInventorySourceRead,
		UpdateContext: resourceInventorySourceUpdate,
		DeleteContext: resourceInventorySourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"inventory_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"overwrite_vars": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				ForceNew: true,
				Optional: true,
			},
			"verbosity": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
			"update_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  0,
				Optional: true,
			},
			"update_on_launch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "scm",
				Optional: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					isTrue := false
					list := []string{"file", "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "tower", "custom"}
					for _, element := range list {
						if element == value {
							isTrue = true
						}
					}
					if !isTrue {
						errs = append(errs, fmt.Errorf("%q must be one of this elements %v, got: %s", key, list, value))
					}
					return
				},
			},
			"source_project_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_path": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
			"source_script": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"credential_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceInventorySourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService

	result, err := awxService.CreateInventorySource(validateInventoryInput(d), map[string]string{})

	if err != nil {
		return DiagsError(InventorySourceResourceName, err)
	}

	d.SetId(getStateID(result.ID))
	return resourceInventorySourceRead(ctx, d, m)

}

func resourceInventorySourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}

	res, err := awxService.UpdateInventorySource(id, validateInventoryInput(d), nil)

	if err != nil {

		return DiagUpdateFail(InventorySourceResourceName, id, err)
	}
	d.SetId(getStateID(res.ID))
	return resourceInventorySourceRead(ctx, d, m)
}

func resourceInventorySourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}
	if _, err := awxService.DeleteInventorySource(id); err != nil {
		return DiagDeleteFail(
			InventorySourceResourceName,
			fmt.Sprintf("inventroy source %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return nil
}

func resourceInventorySourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}
	res, err := awxService.GetInventorySourceByID(id, map[string]string{})
	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}
	setInventorySourceResourceData(d, res)
	return nil
}

//nolint:errcheck,unparam
func setInventorySourceResourceData(d *schema.ResourceData, r *tower.InventorySource) (*schema.ResourceData, diag.Diagnostics) {

	d.Set("name", r.Name)
	d.Set("inventory_id", r.Inventory)
	d.Set("overwrite_vars", r.OverwriteVars)
	d.Set("verbosity", r.Verbosity)
	d.Set("source", r.Source)
	d.Set("source_project_id", r.SourceProject)
	d.Set("source_path", r.SourcePath)
	d.Set("source_script", r.SourceScript)
	d.SetId(getStateID(r.ID))

	return d, nil
}
func validateInventoryInput(d *schema.ResourceData) map[string]interface{} {

	var credential, projectId, sourceScript interface{}

	if d.Get("credential_id").(int) != 0 {
		credential = d.Get("credential_id").(int)
	} else {
		credential = nil
	}
	if d.Get("source_project_id").(int) != 0 {
		projectId = d.Get("source_project_id").(int)
	} else {
		projectId = nil
	}
	if d.Get("source_script").(int) != 0 {
		sourceScript = d.Get("source_script").(int)
	} else {
		sourceScript = nil
	}

	return map[string]interface{}{
		"name":             d.Get("name").(string),
		"inventory":        d.Get("inventory_id").(int),
		"overwrite_vars":   d.Get("overwrite_vars").(bool),
		"verbosity":        d.Get("verbosity").(int),
		"source":           d.Get("source").(string),
		"credential":       credential,
		"source_project":   projectId,
		"update_on_launch": d.Get("update_on_launch").(bool),
		"source_path":      d.Get("source_path").(string),
		"source_script":    sourceScript,
		"overwrite":        d.Get("overwrite").(bool),
	}
}
