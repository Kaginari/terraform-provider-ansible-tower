package tower

import (
	"context"
	"fmt"

	"github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInventoryScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventoryScriptCreate,
		ReadContext:   resourceInventoryScriptRead,
		UpdateContext: resourceInventoryScriptUpdate,
		DeleteContext: resourceInventoryScriptDelete,

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
				Required: true,
			},
			"organization_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceInventoryScriptDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.InventoryScriptsService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventorySourceResourceName, id, err)
	}
	if _, err := awxService.DeleteInventoryScript(id); err != nil {
		return DiagDeleteFail(
			InventoryScriptResourceName,
			fmt.Sprintf("inventroy script %v, got %s ",
				id, err.Error()))
	}
	data.SetId("")
	return nil
}

func resourceInventoryScriptUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.InventoryScriptsService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}

	res, err := awxService.UpdateInventoryScript(id, map[string]interface{}{
		"name":         data.Get("name").(string),
		"description":  data.Get("description").(string),
		"organization": data.Get("organization_id").(int),
		"script":       data.Get("script").(string),
	}, nil)

	if err != nil {

		return DiagUpdateFail(InventorySourceResourceName, id, err)
	}
	data.SetId(getStateID(res.ID))
	return resourceInventoryScriptRead(ctx, data, i)
}

func resourceInventoryScriptCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.InventoryScriptsService

	result, err := awxService.CreateInventoryScript(map[string]interface{}{
		"name":         data.Get("name").(string),
		"description":  data.Get("description").(string),
		"organization": data.Get("organization_id").(int),
		"script":       data.Get("script").(string),
	}, map[string]string{})

	if err != nil {
		return DiagsError(InventoryScriptResourceName, err)
	}

	data.SetId(getStateID(result.ID))
	return resourceInventoryScriptRead(ctx, data, i)
}
func resourceInventoryScriptRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.InventoryScriptsService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}
	res, err := awxService.GetInventoryScriptByID(id, map[string]string{})
	if err != nil {
		return DiagNotFoundFail(InventoryScriptResourceName, id, err)
	}
	setInventoryScriptResourceData(data, res)
	return nil
}

//nolint:errcheck,unparam
func setInventoryScriptResourceData(d *schema.ResourceData, r *tower.InventoryScript) (*schema.ResourceData, diag.Diagnostics) {

	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("organization_id", r.Organization)
	d.Set("script", r.Script)
	d.SetId(getStateID(r.ID))

	return d, nil
}
