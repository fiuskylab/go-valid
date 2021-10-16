package validator

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	{
		want := ""
		got := Min("foo", 1)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldMin, 7)
		got := Min("foo", 7)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Min([]string{"foo", "bar"}, 1)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldMin, 7)
		got := Min([]string{"foo", "bar"}, 7)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
