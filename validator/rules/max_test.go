package validator

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	{
		want := ""
		got := Max("foo", 4)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldLengthMax, 2)
		got := Max("foo", 2)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Max([]string{"foo", "bar"}, 3)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldLengthMax, 1)
		got := Max([]string{"foo", "bar"}, 1)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := ""
		got := Max(20, 30)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(fieldValueMax, 10)
		got := Max(20, 10)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
