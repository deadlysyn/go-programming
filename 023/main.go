package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := x[:5]
	b := x[5:]
	c := x[2:7]
	d := x[1:6]

	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
