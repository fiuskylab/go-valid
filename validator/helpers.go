package validator

import (
	"reflect"
	"strings"
	"unicode"
)

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func toLower(str string) string {
	newStr := ""
	firstUpper := true
	for _, c := range str {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			if firstUpper {
				firstUpper = false
				newStr = newStr + strings.ToLower(string(c))
				continue
			}
			newStr = newStr + "_" + strings.ToLower(string(c))
			continue
		}
		newStr = newStr + string(c)
	}
	return newStr
}
