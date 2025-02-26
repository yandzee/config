package tests

import (
	"reflect"
	"testing"

	"github.com/yandzee/config/str"
)

type SplitTest struct {
	Str      string
	Seps     []string
	Expected []string
}

func runSplitTests(t *testing.T, td []SplitTest) {
	p := &str.StringParser{}

	for i, test := range td {
		result, err := p.Strings(test.Str, test.Seps...)
		if err != nil {
			t.Fatalf("Split test %d has failed: %v\n", i, err.Error())
		}

		if !reflect.DeepEqual(result, test.Expected) {
			t.Fatalf(
				"Split test %d has failed: expected %v (len: %d), got: %v (len: %d)\n",
				i,
				test.Expected,
				len(test.Expected),
				result,
				len(result),
			)
		}
	}
}

func TestStringSplits(t *testing.T) {
	runSplitTests(t, []SplitTest{
		{
			Str:      "",
			Seps:     []string{",", ";", ":"},
			Expected: []string{""},
		},
		{
			Str:      "a",
			Seps:     []string{",", ";", ":"},
			Expected: []string{"a"},
		},
		{
			Str:      "a,b",
			Seps:     []string{"a,b"},
			Expected: []string{"", ""},
		},
		{
			Str:      "a,b",
			Seps:     []string{"a,b", ","},
			Expected: []string{"", ""},
		},
		{
			Str:      "a,b;cde:f",
			Seps:     []string{",", ";", ":"},
			Expected: []string{"a", "b", "cde", "f"},
		},
		{
			Str:      "a:b;cde,f",
			Seps:     []string{",", ";", ":"},
			Expected: []string{"a", "b", "cde", "f"},
		},
		{
			Str:      ":;,:;,,",
			Seps:     []string{",", ",", ","},
			Expected: []string{":;", ":;", "", ""},
		},
		{
			Str:      ":;,:;,,",
			Seps:     []string{",", ":;"},
			Expected: []string{"", "", "", "", "", ""},
		},
	})
}

func BenchmarkStringSplit(b *testing.B) {
	p := &str.StringParser{}
	seps := []string{",", ";", ":"}

	for b.Loop() {
		_, _ = p.Strings("a:b;cde,f", seps...)
	}
}
