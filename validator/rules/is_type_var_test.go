package rules

import (
	"fmt"
	"testing"
)

func TestIsTypeVar(t *testing.T) {
	{
		want := ""
		got := IsTypeVar(int(8), int(256))
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
	{
		want := fmt.Sprintf(isTypeVar, "int64", "int8")
		got := IsTypeVar(int8(8), int64(256))
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
}
