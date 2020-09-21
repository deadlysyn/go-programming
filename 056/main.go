package main

import "fmt"

func main() {
	// // error: send only channel can't receive on line 13
	// // cs := make(chan<- int)
	// cs := make(chan int)
	// go func() {
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs)
	// fmt.Printf("------\n")
	// fmt.Printf("cs\t%T\n", cs)

	// error: receive only channel can't send on line 20
	// cr := make(<-chan int)
	cr := make(chan int)
	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
