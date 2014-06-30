package main

import (
	"go/ast"
	"go/parser"
	"go/token"

	. "gist.github.com/5259939.git"
	"github.com/shurcooL/go-goon"
)

func foo(bar int) int { return bar * 2 }

func main() {
	//goon.Dump(ast.NewObj(0, "name"));return

	fset := token.NewFileSet()
	if file, err := parser.ParseFile(fset, GetThisGoSourceFilepath(), nil, 0); nil == err {
		for _, d := range file.Decls {
			if f, ok := d.(*ast.FuncDecl); ok {
				goon.Dump(f)
				break
			}
		}
	}
}
