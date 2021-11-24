package tower

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	InventoryResourceName         = "Ansible tower inventory"
	OrganisationResourceName      = "Ansible tower Organisation"
	InventorySourceResourceName   = "Ansible tower inventory source"
	InventoryScriptResourceName   = "Ansible tower inventory script"
	CredentialSCMResourceName     = "Ansible tower credential scm"
	CredentialMachineResourceName = "Ansible tower credential machine"
	CredentialTypeResourceName 	  = "Ansible tower credential Type"
)
const (
	InfoColor    = "\033[1;34m[Info] : %s\033[0m"
	NoticeColor  = "\033[1;36m[Notice] : %s\033[0m"
	WarningColor = "\033[1;33m[Warn] : %s\033[0m"
	ErrorColor   = "\033[1;31m[Error] : %s\033[0m"
	DebugColor   = "\033[0;36m[Debug] : %s\033[0m"
)

func SprintError(message string) string {
	return fmt.Sprintf(ErrorColor, message)
}
func SprintWarning(message string) string {
	return fmt.Sprintf(WarningColor, message)
}
func SprintInfo(message string) string {
	return fmt.Sprintf(InfoColor, message)
}
func SprintNotice(message string) string {
	return fmt.Sprintf(NoticeColor, message)
}
func SprintDebug(message string) string {
	return fmt.Sprintf(DebugColor, message)
}
func DiagsError(resource string, err error) diag.Diagnostics {
	return DiagnosticsMessage(
		SprintWarning(fmt.Sprintf("Unable to create %s", resource)),
		SprintError(fmt.Sprintf("Error %s", err.Error())),
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
		SprintWarning(fmt.Sprintf("Unable to fetch %s", resource)),
		SprintWarning(fmt.Sprintf("Unable to load %s with id %d: got %s", resource, id, err.Error())),
	)
}

func DiagUpdateFail(resource string, id int, err error) diag.Diagnostics {
	return DiagnosticsMessage(
		SprintWarning(fmt.Sprintf("Unable to update %s", resource)),
		SprintError(fmt.Sprintf("Unable to update %s with id %d: got %s", resource, id, err.Error())),
	)
}
func DiagDeleteFail(resource, details string) diag.Diagnostics {
	return DiagnosticsMessage(
		SprintWarning(fmt.Sprintf("%s delete faild", resource)),
		SprintError(fmt.Sprintf("Fail to delete %s, %s", resource, details)),
	)
}
func DiagCreateFail(resource string, details string) diag.Diagnostics {
	return DiagnosticsMessage(
		SprintWarning(fmt.Sprintf("Unable to create %s", resource)),
		SprintError(details),
	)
}
