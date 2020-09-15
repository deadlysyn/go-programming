package main

import "fmt"

func receiver(f func(s string)) {
	fmt.Println("doing some work")
	f("all done")
}

func sent(s string) {
	fmt.Println(s)
}

func main() {
	receiver(sent)
}
