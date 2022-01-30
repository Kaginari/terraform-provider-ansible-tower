package tower

import (
	"context"
	"fmt"
	tower "github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceCredentialType() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialTypeCreate,
		ReadContext:   resourceCredentialTypeRead,
		UpdateContext: resourceCredentialTypeUpdate,
		DeleteContext: resourceCredentialTypeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud",
			},
			"input": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				MaxItems: 5,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"label": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"help_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"multiline": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}
func resourceCredentialTypeRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.CredentialTypeService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(CredentialTypeResourceName, err)
	}
	r, err := awxService.GetCredentialsTypesByID(id, map[string]string{})

	if err != nil {
		return DiagNotFoundFail(CredentialTypeResourceName, id, err)
	}
	data = setCredentialTypeResourceData(data, r)
	return nil
}

func resourceCredentialTypeDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.CredentialTypeService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(CredentialTypeResourceName, err)
	}
	if err := awxService.DeleteCredentialsTypesByID(id, map[string]string{}); err != nil {
		return DiagDeleteFail(
			CredentialTypeResourceName,
			fmt.Sprintf(
				"%s %v, got %s ",
				CredentialTypeResourceName, id, err.Error(),
			),
		)
	}
	data.SetId("")
	return nil
}

func resourceCredentialTypeUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	awxService := client.CredentialTypeService
	stateID := data.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(CredentialTypeResourceName, id, err)
	}
	var inputs []Fields
	list := data.Get("input").(*schema.Set).List()
	err = mapstructure.Decode(list, &inputs)
	if err != nil {
		return DiagsError(CredentialTypeResourceName, err)
	}
	credentialInputList := CreateCredentialInputs(inputs)
	_, err = awxService.UpdateCredentialsTypesByID(id, map[string]interface{}{
		"name":   data.Get("name").(string),
		"kind":   data.Get("kind").(string),
		"inputs": credentialInputList,
	}, nil)

	if err != nil {
		return DiagUpdateFail(CredentialTypeResourceName, id, err)
	}

	return resourceCredentialTypeRead(ctx, data, i)
}

func resourceCredentialTypeCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*tower.AWX)
	towerService := client.CredentialTypeService
	var inputs []Fields
	list := data.Get("input").(*schema.Set).List()
	err := mapstructure.Decode(list, &inputs)
	if err != nil {
		return DiagsError(CredentialTypeResourceName, err)
	}
	credentialInputList := CreateCredentialInputs(inputs)
	result, err := towerService.CreateCredentialsTypes(map[string]interface{}{
		"name":   data.Get("name").(string),
		"kind":   data.Get("kind").(string),
		"inputs": credentialInputList,
	}, map[string]string{})
	if err != nil {
		return DiagsError(CredentialTypeResourceName, err)
	}
	data.SetId(getStateID(result.ID))
	return resourceCredentialTypeRead(ctx, data, i)
}

//nolint:errcheck
func setCredentialTypeResourceData(d *schema.ResourceData, r *tower.CredentialType) *schema.ResourceData {
	d.Set("name", r.Name)
	d.Set("kind", r.Kind)
	d.SetId(getStateID(r.ID))
	return d
}
