package main

import (
	"encoding/json"
	"fmt"
)

type TestStruct struct {
	BoolVar      bool        `json:"bool_var,omitempty"`
	IntVar       int         `json:"int_var,omitempty"`
	StringVar    string      `json:"string_var,omitempty"`
	InterfaceVar interface{} `json:"interface,omitempty"`
}

func main() {
	InterfaceVarIsInt := TestStruct{
		BoolVar:      false,
		InterfaceVar: 0,
	}
	res, _ := json.Marshal(InterfaceVarIsInt)
	fmt.Println(string(res))

	var test TestStruct

	err := json.Unmarshal(res, &test)
	if err != nil {

	}
	fmt.Println(test.IntVar == 0)
	//InterfaceVarIsBool := TestStruct{
	//	BoolVar:      false,
	//	IntVar:       0,
	//	InterfaceVar: false,
	//	StringVar: "",
	//}
	//res, _ = json.Marshal(InterfaceVarIsBool)
	//fmt.Println(string(res))
}
