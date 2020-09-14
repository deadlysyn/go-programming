package main

import "fmt"

func main() {
	s := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   31,
	}

	fmt.Println(s)
	fmt.Println(s.age)
}
