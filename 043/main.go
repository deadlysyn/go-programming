package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.age++
}

func main() {
	p := person{
		first: "Rob",
		last:  "Pike",
		age:   42,
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
