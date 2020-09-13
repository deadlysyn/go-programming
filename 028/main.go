package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james": {
			"Shaken, not stirred",
			"Martinis",
			"Women",
		},
		"moneypenny_miss": {
			"James Bond",
			"Literature",
			"Computer Science",
		},
		"no_dr": {
			"Being evil",
			"Ice cream",
			"Sunsets",
		},
	}

	for k, v := range m {
		fmt.Printf("%v:\n", k)
		for i, s := range v {
			fmt.Printf("\t%v\t%v\n", i, s)
		}
	}

}
