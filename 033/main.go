package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	t1 := foo(x...)
	fmt.Println(t1)

	t2 := bar(x)
	fmt.Println(t2)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
