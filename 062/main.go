package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never",
		},
	}

	// don't throw away the error, check them where they happen!
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
