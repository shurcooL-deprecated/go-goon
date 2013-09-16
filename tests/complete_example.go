package main

import "github.com/shurcooL/go-goon"

func main() {
	goon.Dump(map[string]int64{
		"x": 1,
		"y": 4,
		"z": 7,
	})

	goon.Dump([]int32{1, 5, 8})

	{
		x := (*string)(nil)
		goon.Dump(x, nil)
	}

	goon.Dump([]byte("foodboohbingbongstrike123"))

	goon.Dump(uintptr(0), uintptr(123))

	{
		f := func() { println("This is a func.") }

		goon.Dump(f)

		f2 := func(a int, b int) int {
			c := a + b
			return c
		}

		goon.Dump(f2)

		unexportedFuncStruct := struct {
			unexportedFunc func() string
		}{func() string { return "This is the source of an unexported struct field." }}

		goon.Dump(unexportedFuncStruct)
	}
}
