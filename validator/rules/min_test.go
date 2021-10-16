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
		want := fmt.Sprintf(fieldLengthMin, 7)
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
		want := fmt.Sprintf(fieldLengthMin, 7)
		got := Min([]string{"foo", "bar"}, 7)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Min(20, 10)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldValueMin, 10)
		got := Min(7, 10)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
