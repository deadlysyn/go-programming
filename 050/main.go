package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello, my name is", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"James Bond"}

	// we can pass type *person into saySomething
	saySomething(&p)

	// we can NOT pass type person into saySomething
	// (since it has a pointer receiver)
	// saySomething(p)
}
