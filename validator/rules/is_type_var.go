package rules

import (
	"fmt"
	"reflect"
)

const (
	isTypeVar = `expected "%s", received "%s"`
)

func IsTypeVar(v interface{}, t interface{}) string {
	typeV := reflect.TypeOf(v)
	typeT := reflect.TypeOf(t)

	fmt.Println(typeV.Name())

	if typeT != typeV {
		return fmt.Sprintf(isTypeVar, typeT.Name(), typeV.Name())
	}

	return ""
}
