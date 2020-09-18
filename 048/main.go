package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Yyyyyyyy, meeeeee",
			"Alpha beta gamma max",
			"Shaken, not stirred",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"Uh oh",
			"Why me god",
			"ABCDEFGHIJKLMNOP",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Where the hell is James?!?",
			"123 and 321~",
			"Noooooooooooooooooooo",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("Start:", users)

	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	fmt.Println("By age:", users)

	sort.Slice(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	fmt.Println("By name:", users)

	for k := range users {
		fmt.Println("Sayings for", users[k].First, users[k].Last)
		sort.Slice(users[k].Sayings, func(i, j int) bool { return users[k].Sayings[i] < users[k].Sayings[j] })
		for _, v := range users[k].Sayings {
			fmt.Printf("\t%v\n", v)
		}
	}

}
