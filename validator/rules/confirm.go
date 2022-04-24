package rules

import (
	"fmt"
	"reflect"
)

const (
	confirmMessage = `"%s" don't match with "%s_confirmation"`
)

func Confirm(i, j interface{}, fieldName string) string {
	if reflect.DeepEqual(i, j) {
		return ""
	}
	return fmt.Sprintf(confirmMessage, fieldName, fieldName)
}
