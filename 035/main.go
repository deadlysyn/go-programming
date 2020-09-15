package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello from %v, I am %v\n", p.first, p.age)
}

func main() {
	annie := person{
		first: "Annie",
		last:  "Yuan",
		age:   21,
	}
	annie.speak()
}
