package main

import "fmt"

func main() {
	func(name string, greeting string) {
		fmt.Printf("%v, %v...", greeting, name)
	}("James", "Hello")
}
