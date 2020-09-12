package main

import "fmt"

func main() {
	year := 1979
	for {
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}
