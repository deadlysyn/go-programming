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

	people := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	fmt.Println(people[p1.first])
	for i, v := range people[p1.first].iceCreams {
		fmt.Println(i, v)
	}

	fmt.Println(people[p2.first])
	for i, v := range people[p2.first].iceCreams {
		fmt.Println(i, v)
	}
}
