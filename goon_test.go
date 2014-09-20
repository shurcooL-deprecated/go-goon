package goon_test

import (
	"testing"

	"github.com/shurcooL/go-goon"
)

func Example() {
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
}

func TestMap(t *testing.T) {
	got := goon.Sdump(map[string]int64{
		"x": 1,
		"z": 7,
	})

	expected := []string{`(map[string]int64)(map[string]int64{
	(string)("x"): (int64)(1),
	(string)("z"): (int64)(7),
})
`, `(map[string]int64)(map[string]int64{
	(string)("z"): (int64)(7),
	(string)("x"): (int64)(1),
})
`}

	if got != expected[0] && got != expected[1] {
		t.Errorf("got %s", got)
	}
}

func ExampleComplete() {
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

	// Output:
	//([]int32)([]int32{
	//	(int32)(1),
	//	(int32)(5),
	//	(int32)(8),
	//})
	//(*string)(nil)
	//(interface{})(nil)
	//([]uint8)([]uint8{
	//	(uint8)(102),
	//	(uint8)(111),
	//	(uint8)(111),
	//	(uint8)(100),
	//	(uint8)(98),
	//	(uint8)(111),
	//	(uint8)(111),
	//	(uint8)(104),
	//	(uint8)(98),
	//	(uint8)(105),
	//	(uint8)(110),
	//	(uint8)(103),
	//	(uint8)(98),
	//	(uint8)(111),
	//	(uint8)(110),
	//	(uint8)(103),
	//	(uint8)(115),
	//	(uint8)(116),
	//	(uint8)(114),
	//	(uint8)(105),
	//	(uint8)(107),
	//	(uint8)(101),
	//	(uint8)(49),
	//	(uint8)(50),
	//	(uint8)(51),
	//})
	//(uintptr)(nil)
	//(uintptr)(0x7b)
	//(func())(func() { println("This is a func.") })
	//(func(int, int) int)(func(a int, b int) int {
	//	c := a + b
	//	return c
	//})
	//(struct{ unexportedFunc func() string })(struct{ unexportedFunc func() string }{
	//	unexportedFunc: (func() string)(func() string { return "This is the source of an unexported struct field." }),
	//})
	//
}
