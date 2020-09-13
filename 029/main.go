package main

import "fmt"

type person struct {
	first     string
	last      string
	iceCreams []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		iceCreams: []string{
			"chocolate",
			"pistashio",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		iceCreams: []string{
			"strawberry",
			"hazelnut",
		},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.iceCreams {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println(p2.first, p2.last)
	for _, v := range p2.iceCreams {
		fmt.Printf("\t%v\n", v)
	}
}
