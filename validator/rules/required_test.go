package rules

import (
	"testing"
)

func TestRequired(t *testing.T) {
	{
		got := Required("")
		want := fieldRequired
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
	{
		got := Required("foo")
		want := ""
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
	{
		got := Required([]int{})
		want := fieldRequired
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
	{
		got := Required([]int{1, 2, 3})
		want := ""
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
	{
		got := Required([]string{})
		want := fieldRequired
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
	{
		got := Required([]string{"foo", "bar"})
		want := ""
		if got != want {
			t.Errorf(`Got: "%s", want: "%s"`, got, want)
		}
	}
}
