package main

import "fmt"

func main() {
	x := [][]string{
		{
			"James",
			"Bond",
			"Shaken, not stirred",
		},
		{
			"Miss",
			"Moneypenny",
			"Helloooooo, James.",
		},
	}

	for _, v := range x {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
