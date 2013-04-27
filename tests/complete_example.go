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
}