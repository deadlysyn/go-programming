package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("calling foo")
	r := Foo()
	fmt.Println(r)

	fmt.Println("calling bar")
	Bar(100000000)
	fmt.Println("done")
}

func Foo() []string {
	return []string{
		"foo",
		"foo",
		"bar",
	}
}

func Bar(x int) {
	for i := 0; i < x; i++ {
		_ = math.Sqrt(float64(i))
	}
}
