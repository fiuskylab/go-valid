package rules

import (
	"fmt"
	"regexp"
)

const (
	regexNotMatch = `The string "%s" don't match with the regex "%s"`
)

func Regex(v interface{}, expr string) string {
	rgx, err := regexp.Compile(expr)

	if err != nil {
		return err.Error()
	}

	switch v.(type) {
	case string:
		if !rgx.MatchString(v.(string)) {
			return fmt.Sprintf(regexNotMatch, v.(string), expr)
		}
		return ""
	}

	return ""
}
