package goon_test

import "testing"

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	. "github.com/shurcooL/go/gists/gist5286084"
)

func Test(t *testing.T) {
	err := os.Chdir("./tests/")
	CheckError(err)
	files, err := ioutil.ReadDir("./")
	CheckError(err)

	cmds := []func(string) string{
		func(filename string) string { return "go run \"" + filename + "\" > \"" + filename + ".out\"" },
		func(filename string) string { return "git diff --no-ext-diff -- \"" + filename + ".out\"" },
	}

	for _, cmd := range cmds {
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
				filename := file.Name()
				cmdString := cmd(filename)

				out, err := exec.Command("bash", "-c", cmdString).CombinedOutput()
				if nil != err || 0 != len(out) {
					t.Errorf("Failed `%s` with err %v and output %q.", cmdString, err, string(out))
				}
			}
		}
	}
}
