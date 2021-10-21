package tower

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)
const (
	InventoryResourceName       = "Ansible tower inventory"
	OrganisationResourceName	= "Ansible tower Organisation"
)
func DiagsError(resource string, err error) diag.Diagnostics {
	return DiagnosticsMessage(
		fmt.Sprintf("Unable to create %s", resource),
		fmt.Sprintf("Error %s", err.Error()),
	)
}

func DiagnosticsMessage(summary string, details string) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  summary,
		Detail:   details,
	})
	return diags
}
func DiagNotFoundFail(resource string, id int, err error) diag.Diagnostics {
	return DiagnosticsMessage(
		fmt.Sprintf("Unable to fetch %s", resource),
		fmt.Sprintf("Unable to load %s with id %d: got %s", resource, id, err.Error()),
	)
}

func DiagUpdateFail(resource string, id int, err error) diag.Diagnostics {
	return DiagnosticsMessage(
		fmt.Sprintf("Unable to update %s", resource),
		fmt.Sprintf("Unable to update %s with id %d: got %s", resource, id, err.Error()),
	)
}
func DiagDeleteFail(resource, details string) diag.Diagnostics {
	return DiagnosticsMessage(
		fmt.Sprintf("%s delete faild", resource),
		fmt.Sprintf("Fail to delete %s, %s", resource, details),
	)
}
