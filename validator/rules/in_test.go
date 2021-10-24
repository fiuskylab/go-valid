package rules

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {
	{
		want := ""
		got := In("foo", []interface{}{1, 2, 3.5, "bar", "foo"})
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		slice := []interface{}{1, 2, 3.5, "bar", "foo"}
		want := fmt.Sprintf("value not in: %s", slice)
		got := In("not_in_slice", slice)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
