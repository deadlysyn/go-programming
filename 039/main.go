package main

import "fmt"

func returner() func() {
	return func() {
		fmt.Println("returned")
	}
}

func main() {
	f := returner()
	f()
}
