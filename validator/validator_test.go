package validator

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	String   string
	Int      int
	SliceStr []string
}

func TestCheck(t *testing.T) {
	ts := TestStruct{
		String:   "foo",
		Int:      10,
		SliceStr: []string{"foo", "bar"},
	}

	rules := map[string][]string{
		"string":   {"required", "min:2", "max:10"},
		"int":      {"required", "min:2", "max:7"},
		"slicestr": {"required", "min:3", "max:10"},
	}

	{
		want := Result{}
		got, gotErr := Check(ts, rules)

		if !reflect.DeepEqual(want, got) {
			t.Errorf(`want "%T", got "%T"`, want, got)
		}
		if gotErr != nil {
			t.Errorf(`Error must be nil, received %s`, gotErr.Error())
		}
	}
}
