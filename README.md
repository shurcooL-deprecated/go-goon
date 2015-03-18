goon [![Build Status](https://travis-ci.org/shurcooL/go-goon.svg?branch=master)](https://travis-ci.org/shurcooL/go-goon)
=======

A deep pretty printer with Go-like notation. It implements the [goon](https://github.com/shurcooL/goon) specification.

Installation
------------

```bash
go get -u github.com/shurcooL/go-goon
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

Attribution
-----------

go-goon source was based on the existing source of [go-spew](https://github.com/davecgh/go-spew) by Dave Collins.

License
-------

- [MIT License](http://opensource.org/licenses/mit-license.php)
