go-goon
=======

Go-goon is a WIP Go implementation of [goon](https://github.com/shurcooL/goon).

Go-goon currently seems to work for _more_ of the inputs than previously, but could potentially panic and produce wrong results for some inputs. It doesn't reproduce circular structures accurately. There are some things hardcoded and unfinished, so at this time it's not production ready. I plan to fix problems and add missing functionality as I use it in other situations and notice it break.

Installation
------------

```bash
$ go get -u github.com/shurcooL/go-goon
```

Example Output
--------------
[small_example.go](tests/small_example.go) produces:
```go
(Lang)(Lang{
	Name: (string)("Go"),
	Year: (int)(2009),
	URL:  (string)("http"),
	Inner: (*Inner)(&Inner{
		Field1: (string)("Secret!"),
		Field2: (int)(0),
	}),
})

```

[complete_example.go](tests/complete_example.go) produces:
```go
(map[string]int64)(map[string]int64{
	(string)("x"): (int64)(1),
	(string)("y"): (int64)(4),
	(string)("z"): (int64)(7),
})
([]int32)([]int32{
	(int32)(1),
	(int32)(5),
	(int32)(8),
})
(*string)(nil)
(interface{})(nil)
([]uint8)([]uint8{
	(uint8)(102),
	(uint8)(111),
	(uint8)(111),
	(uint8)(100),
	(uint8)(98),
	(uint8)(111),
	(uint8)(111),
	(uint8)(104),
	(uint8)(98),
	(uint8)(105),
	(uint8)(110),
	(uint8)(103),
	(uint8)(98),
	(uint8)(111),
	(uint8)(110),
	(uint8)(103),
	(uint8)(115),
	(uint8)(116),
	(uint8)(114),
	(uint8)(105),
	(uint8)(107),
	(uint8)(101),
	(uint8)(49),
	(uint8)(50),
	(uint8)(51),
})
(uintptr)(nil)
(uintptr)(0x7b)
(func())(func() { println("This is a func.") })
(func(int, int) int)(func(a int, b int) int {
	c := a + b
	return c
})
(struct{ unexportedFunc func() string })(struct{ unexportedFunc func() string }{
	unexportedFunc: (func() string)(func() string { return "This is the source of an unexported struct field." }),
})

```

[large_example.go](tests/large_example.go) produces:
```go
(*ast.FuncDecl)(&ast.FuncDecl{
	Doc:  (*ast.CommentGroup)(nil),
	Recv: (*ast.FieldList)(nil),
	Name: (*ast.Ident)(&ast.Ident{
		NamePos: (token.Pos)(131),
		Name:    (string)("foo"),
		Obj: (*ast.Object)(&ast.Object{
			Kind: (ast.ObjKind)(5),
			Name: (string)("foo"),
			Decl: (*ast.FuncDecl)(nil /*<already shown>*/),
			Data: (interface{})(nil),
			Type: (interface{})(nil),
		}),
	}),
	Type: (*ast.FuncType)(&ast.FuncType{
		Func: (token.Pos)(126),
		Params: (*ast.FieldList)(&ast.FieldList{
			Opening: (token.Pos)(134),
			List: ([]*ast.Field)([]*ast.Field{
				(*ast.Field)(&ast.Field{
					Doc: (*ast.CommentGroup)(nil),
					Names: ([]*ast.Ident)([]*ast.Ident{
						(*ast.Ident)(&ast.Ident{
							NamePos: (token.Pos)(135),
							Name:    (string)("bar"),
							Obj: (*ast.Object)(&ast.Object{
								Kind: (ast.ObjKind)(4),
								Name: (string)("bar"),
								Decl: (*ast.Field)(nil /*<already shown>*/),
								Data: (interface{})(nil),
								Type: (interface{})(nil),
							}),
						}),
					}),
					Type: (*ast.Ident)(&ast.Ident{
						NamePos: (token.Pos)(139),
						Name:    (string)("int"),
						Obj:     (*ast.Object)(nil),
					}),
					Tag:     (*ast.BasicLit)(nil),
					Comment: (*ast.CommentGroup)(nil),
				}),
			}),
			Closing: (token.Pos)(142),
		}),
		Results: (*ast.FieldList)(&ast.FieldList{
			Opening: (token.Pos)(0),
			List: ([]*ast.Field)([]*ast.Field{
				(*ast.Field)(&ast.Field{
					Doc:   (*ast.CommentGroup)(nil),
					Names: ([]*ast.Ident)([]*ast.Ident{}),
					Type: (*ast.Ident)(&ast.Ident{
						NamePos: (token.Pos)(144),
						Name:    (string)("int"),
						Obj:     (*ast.Object)(nil),
					}),
					Tag:     (*ast.BasicLit)(nil),
					Comment: (*ast.CommentGroup)(nil),
				}),
			}),
			Closing: (token.Pos)(0),
		}),
	}),
	Body: (*ast.BlockStmt)(&ast.BlockStmt{
		Lbrace: (token.Pos)(148),
		List: ([]ast.Stmt)([]ast.Stmt{
			(*ast.ReturnStmt)(&ast.ReturnStmt{
				Return: (token.Pos)(150),
				Results: ([]ast.Expr)([]ast.Expr{
					(*ast.BinaryExpr)(&ast.BinaryExpr{
						X: (*ast.Ident)(&ast.Ident{
							NamePos: (token.Pos)(157),
							Name:    (string)("bar"),
							Obj: (*ast.Object)(&ast.Object{
								Kind: (ast.ObjKind)(4),
								Name: (string)("bar"),
								Decl: (*ast.Field)(&ast.Field{
									Doc: (*ast.CommentGroup)(nil),
									Names: ([]*ast.Ident)([]*ast.Ident{
										(*ast.Ident)(&ast.Ident{
											NamePos: (token.Pos)(135),
											Name:    (string)("bar"),
											Obj:     (*ast.Object)(nil /*<already shown>*/),
										}),
									}),
									Type: (*ast.Ident)(&ast.Ident{
										NamePos: (token.Pos)(139),
										Name:    (string)("int"),
										Obj:     (*ast.Object)(nil),
									}),
									Tag:     (*ast.BasicLit)(nil),
									Comment: (*ast.CommentGroup)(nil),
								}),
								Data: (interface{})(nil),
								Type: (interface{})(nil),
							}),
						}),
						OpPos: (token.Pos)(161),
						Op:    (token.Token)(14),
						Y: (*ast.BasicLit)(&ast.BasicLit{
							ValuePos: (token.Pos)(163),
							Kind:     (token.Token)(5),
							Value:    (string)("2"),
						}),
					}),
				}),
			}),
		}),
		Rbrace: (token.Pos)(165),
	}),
})

```

Attribution
-----------

Go-goon source was based on the existing source of [go-spew](https://github.com/davecgh/go-spew) by Dave Collins. Thank you so much Dave!

License
-------

- [MIT License](http://opensource.org/licenses/mit-license.php)
