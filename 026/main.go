package main

import "fmt"

func main() {
	states := make([]string, 10, 10)
	fmt.Printf("%T\t%v\t%v\n", states, len(states), cap(states))

	states = []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
	}

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}
