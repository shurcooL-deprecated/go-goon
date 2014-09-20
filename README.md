go-goon
=======

go-goon is a deep pretty printer with Go-like notation. It implements the [goon](https://github.com/shurcooL/goon) specification.

Installation
------------

```bash
go get -u github.com/shurcooL/go-goon
```

Example
-------

```go
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
//(Lang)(Lang{
//	Name: (string)("Go"),
//	Year: (int)(2009),
//	URL:  (string)("http"),
//	Inner: (*Inner)(&Inner{
//		Field1: (string)("Secret!"),
//		Field2: (int)(0),
//	}),
//})
//
```

Attribution
-----------

go-goon source was based on the existing source of [go-spew](https://github.com/davecgh/go-spew) by Dave Collins.

License
-------

- [MIT License](http://opensource.org/licenses/mit-license.php)
