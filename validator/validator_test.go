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

	rules := Rules{
		"string": Rule{
			Required: true,
			Min:      2,
			Max:      10,
		},
		"int": Rule{
			Required: true,
			Min:      2,
			Max:      7,
		},
		"slicestr": Rule{
			Required: true,
			Min:      3,
			Max:      10,
		},
	}

	{
		want := Result{}
		got, gotErr := Check(ts, rules)

		if !reflect.DeepEqual(want, got) {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
		if gotErr != nil {
			t.Errorf(`Error must be nil, received %s`, gotErr.Error())
		}
	}
}
