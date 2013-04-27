package goon_test

import "github.com/shurcooL/go-goon"
import "testing"

import (
	. "gist.github.com/5286084.git"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

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
			Field2: (int)(0),
		},
	}

	want := `(goon_test.Lang)(goon_test.Lang{
	Name: (string)("Go"),
	Year: (int)(2009),
	URL:  (string)("http"),
	Inner: (*goon_test.Inner)(&goon_test.Inner{
		Field1: (string)("Secret!"),
		Field2: (int)(0),
	}),
})
`

	if got := goon.Sdump(x); got != want {
		t.Errorf("goon.Sdump(%#v) = %v, want %v", x, got, want)
	}
}

// TODO: Next step, refactor the common parts out
func TestSecond(t *testing.T) {
	x := []int{
		5,
		3,
	}

	want := `([]int)([]int{
	(int)(5),
	(int)(3),
})
`

	if got := goon.Sdump(x); got != want {
		t.Errorf("goon.Sdump(%#v) = %v, want %v", x, got, want)
	}
}

func TestThird(t *testing.T) {
	x := (*string)(nil)

	want := `(*string)(nil)
(interface{})(nil)
`

	if got := goon.Sdump(x, nil); got != want {
		t.Errorf("goon.Sdump(%#v) = %v, want %v", x, got, want)
	}
}

func TestFourth(t *testing.T) {
	os.Chdir("./tests/")
	files, err := ioutil.ReadDir("./")
	CheckError(err)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			filename := file.Name()
			cmd := "go run \""+filename+"\" > \""+filename+".out\""

			out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
			if nil != err || 0 != len(out) {
				t.Errorf("Failed `%s` with err %v and output %q.", cmd, err, string(out))
			}
		}
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			filename := file.Name()
			cmd := "git diff --no-ext-diff -- \""+filename+".out\""

			out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
			if nil != err || 0 != len(out) {
				t.Errorf("Failed `%s` with err %v and output %q.", cmd, err, string(out))
			}
		}
	}
}
