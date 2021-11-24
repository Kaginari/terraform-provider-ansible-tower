package tower

import (
	"context"
	"crypto/tls"
	"github.com/Kaginari/ansible-tower-sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"tower_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOWER_HOST", "http://127.0.0.1"),
			},
			"tower_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOWER_USERNAME", "admin"),
			},
			"tower_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TOWER_PASSWORD", "password"),
			},
			"ssl_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Disable SSL verification of API calls",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ansible-tower_inventory":          resourceInventory(),
			"ansible-tower_organisation":       resourceOrganization(),
			"ansible-tower_inventory_source":   resourceInventorySource(),
			"ansible-tower_inventory_script":   resourceInventoryScript(),
			"ansible-tower_project":            resourceProject(),
			"ansible-tower_job_template":       resourceJobTemplate(),
			"ansible-tower_credential_scm":     resourceCredentialSCM(),
			"ansible-tower_credential_machine": resourceCredentialMachine(),
			"ansible-tower_credential_type": 	resourceCredentialType(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	hostname := d.Get("tower_host").(string)
	username := d.Get("tower_username").(string)
	password := d.Get("tower_password").(string)

	client := http.DefaultClient
	if d.Get("ssl_verify").(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	var diags diag.Diagnostics
	c, err := tower.NewAWX(hostname, username, password, client)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Authentication error",
			Detail:   "Check Host , Username and Password",
		})
		return nil, diags
	}

	return c, diags
}
