package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("calling foo")
	r := foo()
	fmt.Println(r)

	fmt.Println("calling bar")
	bar(100000000)
	fmt.Println("done")
}

func foo() []string {
	return []string{
		"foo",
		"foo",
		"bar",
	}
}

func bar(x int) {
	for i := 0; i < x; i++ {
		_ = math.Sqrt(float64(i))
	}
}
