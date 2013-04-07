package main

import "github.com/shurcooL/go-goon"

func main() {
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
	goon.Dump(x)

	goon.Dump(map[string]int64{
		"x": 8,
		"z": 8,
		"y": 7,
	})

	goon.Dump([]int32{1, 5, 8})
}