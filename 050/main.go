package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("hello, my name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "James Bond",
		age:  32,
	}

	// we can pass type *person into saySomething
	saySomething(&p)

	// we can NOT pass type person into saySomething
	// saySomething(p)
}
