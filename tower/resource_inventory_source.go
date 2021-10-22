package tower

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tower "github.com/mrcrilly/goawx/client"
)

func resourceInventorySource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventorySourceCreate,
		ReadContext:   resourceInventorySourceRead,
		UpdateContext: resourceInventorySourceUpdate,
		DeleteContext: resourceInventorySourceDelete,

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
			"credential_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceInventorySourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService

	result, err := awxService.CreateInventorySource(map[string]interface{}{
		"name":             d.Get("name").(string),
		"inventory":        d.Get("inventory_id").(int),
		"overwrite_vars":   d.Get("overwrite_vars").(bool),
		"verbosity":        d.Get("verbosity").(int),
		"source":           d.Get("source").(string),
		"credential":       d.Get("credential_id").(int),
		"source_project":   d.Get("source_project_id").(int),
		"update_on_launch": d.Get("update_on_launch").(bool),
		"source_path":      d.Get("source_path").(string),
		"overwrite":        d.Get("overwrite").(bool),
	}, map[string]string{})
	if err != nil {
		return DiagCreateFail(InventorySourceResourceName, err)
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceInventorySourceRead(ctx, d, m)

}

func resourceInventorySourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}

	_, err := awxService.UpdateInventorySource(id, map[string]interface{}{
		"name":             d.Get("name").(string),
		"inventory":        d.Get("inventory_id").(int),
		"overwrite_vars":   d.Get("overwrite_vars").(bool),
		"verbosity":        d.Get("verbosity").(int),
		"source":           d.Get("source").(string),
		"credential":       d.Get("credential_id").(int),
		"source_project":   d.Get("source_project_id").(int),
		"update_on_launch": d.Get("update_on_launch").(bool),
		"source_path":      d.Get("source_path").(string),
		"overwrite":        d.Get("overwrite").(bool),
	}, nil)
	if err != nil {
		return buildDiagUpdateFail(diagElementInventorySourceTitle, id, err)
	}

	return resourceInventorySourceRead(ctx, d, m)
}

func resourceInventorySourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	if _, err := awxService.DeleteInventorySource(id); err != nil {
		return buildDiagDeleteFail(
			"inventroy source",
			fmt.Sprintf("inventroy source %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return nil
}

func resourceInventorySourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tower.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	res, err := awxService.GetInventorySourceByID(id, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail(diagElementInventorySourceTitle, id, err)
	}
	d = setInventorySourceResourceData(d, res)
	return nil
}

func setInventorySourceResourceData(d *schema.ResourceData, r *awx.InventorySource) *schema.ResourceData {
	d.Set("name", r.Name)

	d.Set("inventory_id", r.Inventory)
	d.Set("overwrite_vars", r.OverwriteVars)
	d.Set("verbosity", r.Verbosity)
	d.Set("source", r.Source)
	d.Set("source_project_id", r.SourceProject)
	d.Set("source_path", r.SourcePath)

	return d
}
