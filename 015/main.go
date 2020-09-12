package main

import "fmt"

func main() {
	year := 1979
	for {
		if year <= 2020 {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}
}
