package main

import "fmt"

func main() {
	// error: send only channel can't receive on line 13
	// cs := make(chan<- int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
