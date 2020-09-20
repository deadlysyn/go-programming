package main

import "fmt"

func main() {
	// blocks
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)

	// fix: buffered channel (fragile: how big do you need buffer to be?)
	// c := make(chan int, 1)
	// c <- 42
	// fmt.Println(<-c)

	// fix: anonymous self executing func, idiomatic
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
