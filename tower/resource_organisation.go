package tower

import (
	"context"
	"fmt"
	"github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrganization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganizationsCreate,
		ReadContext:   resourceOrganizationsRead,
		UpdateContext: resourceOrganizationsUpdate,
		DeleteContext: resourceOrganizationsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
			"max_hosts": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Maximum number of hosts allowed to be managed by this organization",
			},
			"custom_virtualenv": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Local absolute file path containing a custom Python virtualenv to use",
			},
		},
	}
}

func resourceOrganizationsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.OrganizationsService

	result, err := awxService.CreateOrganization(map[string]interface{}{
		"name":              d.Get("name").(string),
		"description":       d.Get("description").(string),
		"max_hosts":         d.Get("max_hosts").(int),
		"custom_virtualenv": d.Get("custom_virtualenv").(string),
	}, map[string]string{})
	if err != nil {
		return DiagsError(OrganisationResourceName, err)
	}
	d.SetId(getStateID(result.ID))
	return resourceOrganizationsRead(ctx, d, m)
}

func resourceOrganizationsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*tower.AWX)
	awxService := client.OrganizationsService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(OrganisationResourceName, id, err)
	}
	params := make(map[string]string)

	_, err = awxService.GetOrganizationsByID(id, params)
	if err != nil {
		return DiagNotFoundFail("Organizations", id, err)
	}

	_, err = awxService.UpdateOrganization(id, map[string]interface{}{
		"name":              d.Get("name").(string),
		"description":       d.Get("description").(string),
		"max_hosts":         d.Get("max_hosts").(int),
		"custom_virtualenv": d.Get("custom_virtualenv").(string),
	}, map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Organizations",
			Detail:   fmt.Sprintf("Organizations with name %s faild to update %s", d.Get("name").(string), err.Error()),
		})
		return diags
	}

	return resourceOrganizationsRead(ctx, d, m)
}

func resourceOrganizationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*tower.AWX)
	awxService := client.OrganizationsService

	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(OrganisationResourceName, id, err)
	}
	res, err := awxService.GetOrganizationsByID(id, make(map[string]string))
	if err != nil {
		return DiagNotFoundFail("Organization", id, err)

	}
	d = setOrganizationsResourceData(d, res)
	return diags
}

func resourceOrganizationsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	digMessagePart := "Organization"
	client := m.(*tower.AWX)
	awxService := client.OrganizationsService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(OrganisationResourceName, id, err)
	}

	if _, err := awxService.DeleteOrganization(id); err != nil {
		return DiagDeleteFail(digMessagePart, fmt.Sprintf("OrganizationID %v, got %s ", id, err.Error()))
	}
	d.SetId("")
	return diags
}

//nolint:errcheck
func setOrganizationsResourceData(d *schema.ResourceData, r *tower.Organizations) *schema.ResourceData {

	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("max_hosts", r.MaxHosts)
	d.Set("custom_virtualenv", r.CustomVirtualenv)
	d.SetId(getStateID(r.ID))
	return d
}
