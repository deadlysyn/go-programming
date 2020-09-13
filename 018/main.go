package main

import "fmt"

func main() {
	x := 42
	if x > 50 {
		fmt.Println("Won't print")
	} else if x < 40 {
		fmt.Println("Still won't print")
	} else {
		fmt.Println("Finally prints")
	}
}
