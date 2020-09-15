package main

import "fmt"

func main() {
	f := func(x, y int) {
		fmt.Println(x * y)
	}

	f(2, 4)
}
