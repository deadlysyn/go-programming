package main

import "fmt"

func main() {
	defer hello()
	goodbye()
}

func hello() {
	fmt.Println("Hello there")
}

func goodbye() {
	fmt.Println("Buhbye")
}
