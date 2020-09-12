package main

import "fmt"

func main() {
	// print out ascii chars from 33-122
	for i := 33; i <= 122; i++ {
		fmt.Printf("i is %d, ASCII %c\n", i, i)
	}
}
