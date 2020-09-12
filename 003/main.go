package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// print default zero values
	s := fmt.Sprintf("'%v', '%v', '%v'\n", x, y, z)
	fmt.Println(s)
}
