package rules

import (
	"fmt"
	"testing"
)

func TestConfirm(t *testing.T) {
	{
		want := ""
		got := Confirm("foo", "foo", "field_foo")
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
	{
		want := fmt.Sprintf(confirmMessage, "field_foo", "field_foo")
		got := Confirm("foo", "bar", "field_foo")
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
}
