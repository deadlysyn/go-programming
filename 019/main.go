package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("oh yeah")
	case false:
		fmt.Println("oh no")
	default:
		fmt.Println("default")
	}
}
