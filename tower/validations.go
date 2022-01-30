package tower

import (
	"fmt"
)

//inventorySourceValidations
//nolint:deadcode,unused
func inventorySourceValidations(source string, project int) error {
	switch v := source; v {
	case "scm":
		if project == 0 {
			return fmt.Errorf("Your source type is \"%s\" need to specify project id ", source)
		}
	case "custom":
		return nil
		// TODO SUPPORT ALL KIND OF SOURCES
		//case "file":
		//	return nil
		//case "ec2":
		//	return nil
		//case "gce":
		//	return nil
		//case "azure_rm":
		//	return nil
		//case "vmware":
		//	return nil
		//case "satellite6":
		//	return nil
		//case "openstack":
		//	return nil
		//case "rhv":
		//	return nil
		//case "tower":
		//	return nil
	}
	return nil
}
