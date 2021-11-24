package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {

}

func test2() {
	a := "test inventory"
	id := 3
	noSpaceString := strings.ReplaceAll(a, " ", "")
	str := noSpaceString + "." + strconv.Itoa(id)
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encoded)
}
func test() {
	var list []Variable
	var a Variable
	a.Value = "testvalue1"
	a.Key = "keytest"
	var b Variable
	b.Value = "testvalue2"
	b.Key = "keytest2"
	list = append(list, a)
	list = append(list, b)

	var result bytes.Buffer

	result.WriteString("{")
	for index, ele := range list {
		if index != 0 {
			result.WriteString(",")
		}
		result.WriteString("\"")
		result.WriteString(ele.Key)
		result.WriteString("\"")
		result.WriteString(":")
		result.WriteString("\"")
		result.WriteString(ele.Value)
		result.WriteString("\"")
	}
	result.WriteString("}")

	fmt.Println(result.String())
}
