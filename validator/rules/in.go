package rules

import "fmt"

func In(v interface{}, in []interface{}) string {
	for _, i := range in {
		if v == i {
			return ""
		}
	}

	return fmt.Sprintf("value not in: %s", in)
}
