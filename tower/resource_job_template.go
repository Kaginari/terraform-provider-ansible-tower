package tower

import (
	"context"
	"fmt"
	tower "github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

func resourceJobTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJobTemplateCreate,
		ReadContext:   resourceJobTemplateRead,
		UpdateContext: resourceJobTemplateUpdate,
		DeleteContext: resourceJobTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"job_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "One of: run, check",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					isTrue := false
					list := []string{"run", "check"}
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
			"inventory_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"playbook": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"forks": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"limit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"verbosity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "One of 0,1,2,3,4,5",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(int)
					isTrue := false
					list := []int{0, 1, 2, 3, 4, 5}
					for _, element := range list {
						if element == value {
							isTrue = true
						}
					}
					if !isTrue {
						errs = append(errs, fmt.Errorf("%q must be one of this elements %v, got: %d", key, list, value))
					}
					return
				},
			},
			"extra_vars": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"job_tags": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"force_handlers": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"skip_tags": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"start_at_task": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"use_fact_cache": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"host_config_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"ask_diff_mode_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_limit_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_tags_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_verbosity_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_inventory_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_variables_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_credential_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"survey_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"become_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"diff_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ask_skip_tags_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_simultaneous": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"custom_virtualenv": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"ask_job_type_on_launch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourceJobTemplateCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*tower.AWX)
	awxService := client.JobTemplateService
	// TODO (depth) change sleep by getting project sync state
	time.Sleep(10 * time.Second)
	result, err := awxService.CreateJobTemplate(validateJobTemplateInput(data), map[string]string{})

	if err != nil {
		log.Printf("Fail to Create Template %v", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create JobTemplate",
			Detail:   fmt.Sprintf("JobTemplate with name %s in the project id %d, faild to create %s", data.Get("name").(string), data.Get("project_id").(int), err.Error()),
		})
		return diags
	}

	data.SetId(getStateID(result.ID))
	return resourceJobTemplateRead(ctx, data, i)
}

func resourceJobTemplateUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*tower.AWX)
	awxService := client.JobTemplateService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}

	params := make(map[string]string)
	_, err = awxService.GetJobTemplateByID(id, params)
	if err != nil {
		return DiagNotFoundFail("job template", id, err)
	}

	_, err = awxService.UpdateJobTemplate(id, validateJobTemplateInput(data), map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update JobTemplate",
			Detail:   fmt.Sprintf("JobTemplate with name %s in the project id %d faild to update %s", data.Get("name").(string), data.Get("project_id").(int), err.Error()),
		})
		return diags
	}

	return resourceJobTemplateRead(ctx, data, i)
}
func resourceJobTemplateDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*tower.AWX)
	awxService := client.JobTemplateService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}
	_, err = awxService.DeleteJobTemplate(id)
	if err != nil {
		return DiagDeleteFail(
			"JobTemplate",
			fmt.Sprintf(
				"JobTemplateID %v, got %s ",
				id, err.Error(),
			),
		)
	}
	data.SetId("")
	return diags
}
func resourceJobTemplateRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*tower.AWX)
	awxService := client.JobTemplateService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}

	res, err := awxService.GetJobTemplateByID(id, make(map[string]string))
	if err != nil {
		return DiagNotFoundFail("job template", id, err)

	}
	data = setJobTemplateResourceData(data, res)
	return diags
}

//nolint:errcheck
func setJobTemplateResourceData(data *schema.ResourceData, r *tower.JobTemplate) *schema.ResourceData {
	data.Set("allow_simultaneous", r.AllowSimultaneous)
	data.Set("ask_credential_on_launch", r.AskCredentialOnLaunch)
	data.Set("ask_job_type_on_launch", r.AskJobTypeOnLaunch)

	data.Set("ask_limit_on_launch", r.AskLimitOnLaunch)
	data.Set("ask_skip_tags_on_launch", r.AskSkipTagsOnLaunch)
	data.Set("ask_tags_on_launch", r.AskTagsOnLaunch)
	data.Set("ask_variables_on_launch", r.AskVariablesOnLaunch)
	data.Set("description", r.Description)
	data.Set("extra_vars", r.ExtraVars)
	data.Set("force_handlers", r.ForceHandlers)
	data.Set("forks", r.Forks)
	data.Set("host_config_key", r.HostConfigKey)
	data.Set("inventory_id", r.Inventory)
	data.Set("job_tags", r.JobTags)
	data.Set("job_type", r.JobType)
	data.Set("diff_mode", r.DiffMode)
	data.Set("custom_virtualenv", r.CustomVirtualenv)
	data.Set("limit", r.Limit)
	data.Set("name", r.Name)
	data.Set("become_enabled", r.BecomeEnabled)
	data.Set("use_fact_cache", r.UseFactCache)
	data.Set("playbook", r.Playbook)
	data.Set("project_id", r.Project)
	data.Set("skip_tags", r.SkipTags)
	data.Set("start_at_task", r.StartAtTask)
	data.Set("survey_enabled", r.SurveyEnabled)
	data.Set("verbosity", r.Verbosity)
	data.SetId(getStateID(r.ID))
	return data
}

func validateJobTemplateInput(data *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"name":                     data.Get("name").(string),
		"description":              data.Get("description").(string),
		"job_type":                 data.Get("job_type").(string),
		"inventory":                data.Get("inventory_id").(int),
		"project":                  data.Get("project_id").(int),
		"playbook":                 data.Get("playbook").(string),
		"forks":                    data.Get("forks").(int),
		"limit":                    data.Get("limit").(string),
		"verbosity":                data.Get("verbosity").(int),
		"extra_vars":               data.Get("extra_vars").(string),
		"job_tags":                 data.Get("job_tags").(string),
		"force_handlers":           data.Get("force_handlers").(bool),
		"skip_tags":                data.Get("skip_tags").(string),
		"start_at_task":            data.Get("start_at_task").(string),
		"timeout":                  data.Get("timeout").(int),
		"use_fact_cache":           data.Get("use_fact_cache").(bool),
		"host_config_key":          data.Get("host_config_key").(string),
		"ask_diff_mode_on_launch":  data.Get("ask_diff_mode_on_launch").(bool),
		"ask_variables_on_launch":  data.Get("ask_variables_on_launch").(bool),
		"ask_limit_on_launch":      data.Get("ask_limit_on_launch").(bool),
		"ask_tags_on_launch":       data.Get("ask_tags_on_launch").(bool),
		"ask_skip_tags_on_launch":  data.Get("ask_skip_tags_on_launch").(bool),
		"ask_job_type_on_launch":   data.Get("ask_job_type_on_launch").(bool),
		"ask_verbosity_on_launch":  data.Get("ask_verbosity_on_launch").(bool),
		"ask_inventory_on_launch":  data.Get("ask_inventory_on_launch").(bool),
		"ask_credential_on_launch": data.Get("ask_credential_on_launch").(bool),
		"survey_enabled":           data.Get("survey_enabled").(bool),
		"become_enabled":           data.Get("become_enabled").(bool),
		"diff_mode":                data.Get("diff_mode").(bool),
		"allow_simultaneous":       data.Get("allow_simultaneous").(bool),
		"custom_virtualenv":        data.Get("custom_virtualenv").(string),
	}

}
