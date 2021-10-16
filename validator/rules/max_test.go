package validator

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	{
		want := ""
		got := Max("foo", 1)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldLengthMax, 7)
		got := Max("foo", 7)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Max([]string{"foo", "bar"}, 1)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldLengthMax, 7)
		got := Max([]string{"foo", "bar"}, 7)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Max(20, 10)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldValueMax, 10)
		got := Max(7, 10)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
