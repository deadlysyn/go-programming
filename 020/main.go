package main

import "fmt"

func main() {
	favSport := "soccer"
	switch favSport {
	case "baseball":
		fmt.Println("nope")
	case "basketball":
		fmt.Println("close")
	case "soccer":
		fmt.Println("YEAH!")
	default:
		fmt.Println("shouldn't get here...")
	}
}
