goon
====

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/go-goon.svg)](https://pkg.go.dev/github.com/shurcooL/go-goon)

Package goon is a deep pretty printer with Go-like notation. It implements the [goon](https://github.com/shurcooL/goon) specification.

**Deprecated:** This package is old, incomplete, low code quality, and now unmaintained.
See [github.com/hexops/valast](https://github.com/hexops/valast) for a newer package that is the closest known direct replacement.
See the [Alternatives](#alternatives) section for other known entries in this problem space.

Installation
------------

```bash
go get github.com/shurcooL/go-goon
```

Examples
--------

```Go
x := Lang{
	Name: "Go",
	Year: 2009,
	URL:  "http",
	Inner: &Inner{
		Field1: "Secret!",
	},
}

goon.Dump(x)

// Output:
// (Lang)(Lang{
// 	Name: (string)("Go"),
// 	Year: (int)(2009),
// 	URL:  (string)("http"),
// 	Inner: (*Inner)(&Inner{
// 		Field1: (string)("Secret!"),
// 		Field2: (int)(0),
// 	}),
// })
```

```Go
items := []int{1, 2, 3}

goon.DumpExpr(len(items))

// Output:
// len(items) = (int)(3)
```

```Go
adderFunc := func(a int, b int) int {
	c := a + b
	return c
}

goon.DumpExpr(adderFunc)

// Output:
// adderFunc = (func(int, int) int)(func(a int, b int) int {
// 	c := a + b
// 	return c
// })
```

Directories
-----------

| Path                                                            | Synopsis                                                                                    |
|-----------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| [bypass](https://pkg.go.dev/github.com/shurcooL/go-goon/bypass) | Package bypass allows bypassing reflect restrictions on accessing unexported struct fields. |

Alternatives
------------

-	[`go-spew`](https://github.com/davecgh/go-spew) - A deep pretty printer for Go data structures to aid in debugging.
-	[`valast`](https://github.com/hexops/valast) - Convert Go values to their AST.
-	[`repr`](https://github.com/alecthomas/repr) - Python's repr() for Go.

Attribution
-----------

go-goon source was based on the existing source of [go-spew](https://github.com/davecgh/go-spew) by [Dave Collins](https://github.com/davecgh).

License
-------

-	[MIT License](LICENSE)
