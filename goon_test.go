package goon_test

import "github.com/shurcooL/go-goon"
import "testing"

func TestFirst(t *testing.T) {
	type Inner struct {
		Field1 string
		Field2 int
	}
	type Lang struct {
		Name  string
		Year  int
		URL   string
		Inner *Inner
	}

	x := Lang{
		Name: "Go",
		Year: 2009,
		URL:  "http",
		Inner: &Inner{
			Field1: "Secret!",
		},
	}

	want := `Lang{
	Name: "Go",
	Year: 2009,
	URL:  "http",
	Inner: &Inner{
		Field1: "Secret!",
	},
}
`

	if got := goon.Sdump(x); got != want {
		t.Errorf("goon.Sdump(%#v) = %v, want %v", x, got, want)
	}
}
