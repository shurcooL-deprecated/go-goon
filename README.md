go-goon
=======

Go-goon is a WIP Go implementation of [goon](https://github.com/shurcooL/goon).

Go-goon currently works for only _SOME_ of the inputs, and likely panics and produces wrong results for many others. There are some things hardcoded and unfinished, so at this time it's not production ready. I plan to fix problems and add missing functionality as I use it in other situations and notice it break.

Installation
------------

```bash
$ go get -u github.com/shurcooL/go-goon
```

Example Output
--------------

```go
Lang{
	Name: "Go",   // (string)
	Year: 2009,   // (int)
	URL:  "http", // (string)
	Inner: &Inner{
		Field1: "Secret!", // (string)
		Field2: 0,         // (int)
	}, // (*main.Inner)
}

```

```go
&ast.FuncDecl{
	Doc:  nil, // (*ast.CommentGroup)
	Recv: nil, // (*ast.FieldList)
	Name: &ast.Ident{
		NamePos: 169,   // (token.Pos)
		Name:    "foo", // (string)
		Obj: &ast.Object{
			Kind: 5,     // (ast.ObjKind)
			Name: "foo", // (string)
			Decl: nil,   /*<already shown>*/ // (*ast.FuncDecl)
			Data: nil,   // (interface {})
			Type: nil,   // (interface {})
		}, // (*ast.Object)
	}, // (*ast.Ident)
	Type: &ast.FuncType{
		Func: 164, // (token.Pos)
		Params: &ast.FieldList{
			Opening: 172, // (token.Pos)
			List: []*ast.Field{
				&ast.Field{
					Doc: nil, // (*ast.CommentGroup)
					Names: []*ast.Ident{
						&ast.Ident{
							NamePos: 173, // (token.Pos)
							Name:    "x", // (string)
							Obj: &ast.Object{
								Kind: 4,   // (ast.ObjKind)
								Name: "x", // (string)
								Decl: nil, /*<already shown>*/ // (*ast.Field)
								Data: nil, // (interface {})
								Type: nil, // (interface {})
							}, // (*ast.Object)
						},
					}, // ([]*ast.Ident)
					Type: &ast.Ident{
						NamePos: 175,   // (token.Pos)
						Name:    "int", // (string)
						Obj:     nil,   // (*ast.Object)
					}, // (*ast.Ident)
					Tag:     nil, // (*ast.BasicLit)
					Comment: nil, // (*ast.CommentGroup)
				},
			}, // ([]*ast.Field)
			Closing: 178, // (token.Pos)
		}, // (*ast.FieldList)
		Results: &ast.FieldList{
			Opening: 0, // (token.Pos)
			List: []*ast.Field{
				&ast.Field{
					Doc:   nil,            // (*ast.CommentGroup)
					Names: []*ast.Ident{}, // ([]*ast.Ident)
					Type: &ast.Ident{
						NamePos: 180,   // (token.Pos)
						Name:    "int", // (string)
						Obj:     nil,   // (*ast.Object)
					}, // (*ast.Ident)
					Tag:     nil, // (*ast.BasicLit)
					Comment: nil, // (*ast.CommentGroup)
				},
			}, // ([]*ast.Field)
			Closing: 0, // (token.Pos)
		}, // (*ast.FieldList)
	}, // (*ast.FuncType)
	Body: &ast.BlockStmt{
		Lbrace: 184, // (token.Pos)
		List: []ast.Stmt{
			&ast.ReturnStmt{
				Return: 186, // (token.Pos)
				Results: []ast.Expr{
					&ast.BinaryExpr{
						X: &ast.Ident{
							NamePos: 193, // (token.Pos)
							Name:    "x", // (string)
							Obj: &ast.Object{
								Kind: 4,   // (ast.ObjKind)
								Name: "x", // (string)
								Decl: &ast.Field{
									Doc: nil, // (*ast.CommentGroup)
									Names: []*ast.Ident{
										&ast.Ident{
											NamePos: 173, // (token.Pos)
											Name:    "x", // (string)
											Obj:     nil, /*<already shown>*/ // (*ast.Object)
										},
									}, // ([]*ast.Ident)
									Type: &ast.Ident{
										NamePos: 175,   // (token.Pos)
										Name:    "int", // (string)
										Obj:     nil,   // (*ast.Object)
									}, // (*ast.Ident)
									Tag:     nil, // (*ast.BasicLit)
									Comment: nil, // (*ast.CommentGroup)
								}, // (*ast.Field)
								Data: nil, // (interface {})
								Type: nil, // (interface {})
							}, // (*ast.Object)
						}, // (*ast.Ident)
						OpPos: 195, // (token.Pos)
						Op:    14,  // (token.Token)
						Y: &ast.BasicLit{
							ValuePos: 197, // (token.Pos)
							Kind:     5,   // (token.Token)
							Value:    "2", // (string)
						}, // (*ast.BasicLit)
					},
				}, // ([]ast.Expr)
			},
		}, // ([]ast.Stmt)
		Rbrace: 199, // (token.Pos)
	}, // (*ast.BlockStmt)
}

```

Attribution
-----------

Go-goon source was based on the existing source of go-spew by Dave Collins. Thank you so much Dave! (I'm not very experienced with dealing with licenses, so if I've done something wrong, please let me know politely and I will politely try to fix it; my intentions are to bring the most value to all humans and I mean no harm.)

License
-------

- [MIT License](http://opensource.org/licenses/mit-license.php)
