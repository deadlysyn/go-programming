package main

import "fmt"

func main() {
	count := closure()
	for i := 0; i < 3; i++ {
		count()
	}
}

func closure() func() {
	x := 0
	return func() {
		x++
		fmt.Println(x)
	}
}
