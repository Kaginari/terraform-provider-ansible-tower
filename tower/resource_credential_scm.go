package tower

import (
	"context"
	"fmt"
	tower "github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCredentialSCM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialSCMCreate,
		ReadContext:   resourceCredentialSCMRead,
		UpdateContext: resourceCredentialSCMUpdate,
		DeleteContext: resourceCredentialSCMDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organisation_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_key_data": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_key_unlock": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceCredentialSCMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	keys := []string{
		"name",
		"description",
		"username",
		"password",
		"ssh_key_data",
		"ssh_public_key_data",
		"ssh_key_unlock",
		"become_method",
		"become_username",
		"become_password",
		"organisation_id",
		"team_id",
		"owner_id",
	}
	if d.HasChanges(keys...) {
		var err error
		newCredential := map[string]interface{}{
			"name":            d.Get("name").(string),
			"description":     d.Get("description").(string),
			"organization":    d.Get("organisation_id").(int),
			"credential_type": 2,
			"inputs": map[string]interface{}{
				"username":       d.Get("username").(string),
				"password":       d.Get("password").(string),
				"ssh_key_data":   d.Get("ssh_key_data").(string),
				"ssh_key_unlock": d.Get("ssh_key_unlock").(string),
			},
		}

		client := m.(*tower.AWX)
		cred, err := client.CredentialsService.CreateCredentials(newCredential, map[string]string{})
		if err != nil {
			return DiagsError(CredentialSCMResourceName, err)
		}

		d.SetId(getStateID(cred.ID))
		resourceCredentialSCMRead(ctx, d, m)
	}
	return diags

}

func resourceCredentialSCMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*tower.AWX)

	stateID := d.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(CredentialSCMResourceName, err)
	}

	cred, err := client.CredentialsService.GetCredentialsByID(id, map[string]string{})

	if err != nil {
		return DiagNotFoundFail(CredentialSCMResourceName, id, err)
	}

	setErr := d.Set("name", cred.Name)
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("description", cred.Description)
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("username", cred.Inputs["username"])
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("password", cred.Inputs["password"])
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("ssh_key_data", cred.Inputs["ssh_key_data"])
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("ssh_key_unlock", cred.Inputs["ssh_key_unlock"])
	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}
	setErr = d.Set("organisation_id", cred.OrganizationID)

	if setErr != nil {
		return DiagsError(CredentialSCMResourceName, setErr)
	}

	return diags
}

func resourceCredentialSCMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(CredentialSCMResourceName, id, err)
	}
	updatedCredential := map[string]interface{}{
		"name":            d.Get("name").(string),
		"description":     d.Get("description").(string),
		"organization":    d.Get("organisation_id").(int),
		"credential_type": 2,
		"inputs": map[string]interface{}{
			"username":       d.Get("username").(string),
			"password":       d.Get("password").(string),
			"ssh_key_data":   d.Get("ssh_key_data").(string),
			"ssh_key_unlock": d.Get("ssh_key_unlock").(string),
		},
	}

	client := m.(*tower.AWX)
	_, err = client.CredentialsService.UpdateCredentialsByID(id, updatedCredential, map[string]string{})
	if err != nil {
		return DiagUpdateFail(CredentialSCMResourceName, id, err)
	}

	return resourceCredentialSCMRead(ctx, d, m)
}
func resourceCredentialSCMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(CredentialSCMResourceName, id, err)
	}
	client := m.(*tower.AWX)
	err = client.CredentialsService.DeleteCredentialsByID(id, map[string]string{})
	if err != nil {
		return DiagDeleteFail(CredentialSCMResourceName, fmt.Sprintf(
			"%s %v, got %s ",
			CredentialSCMResourceName, id, err.Error(),
		))
	}

	return diags
}
