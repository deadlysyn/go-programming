package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	j := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is blah blah"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["Enough is enough","Booyah baby"]}]`
	fmt.Println(j)

	var p []person

	err := json.Unmarshal([]byte(j), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
