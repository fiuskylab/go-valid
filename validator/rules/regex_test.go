package rules

import (
	"fmt"
	"testing"
)

func TestRegex(t *testing.T) {
	{
		want := ""
		got := Regex("something", "[a-z]{1,}")

		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		v := "something"
		r := "[A-Z{1,2}"
		got := Regex(v, r)

		if got == "" {
			t.Errorf(`wanted not nil error, received  %s`, got)
		}
	}
	{
		v := "something"
		r := "[A-Z]{1,2}"
		want := fmt.Sprintf(regexNotMatch, v, r)
		got := Regex(v, r)

		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
