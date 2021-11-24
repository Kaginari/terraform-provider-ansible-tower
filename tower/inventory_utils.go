package tower

import (
	"bytes"
)

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func CreateInventoryVariables(variables []Variable) string {
	var result bytes.Buffer
	result.WriteString("{")
	for index, element := range variables {
		if index != 0 {
			result.WriteString(",")
		}
		result.WriteString(" ")
		result.WriteString(element.Key)
		result.WriteString(" ")
		result.WriteString(":")
		result.WriteString(" ")
		result.WriteString(element.Value)
		result.WriteString(" ")
	}
	result.WriteString("}")
	return result.String()
}
