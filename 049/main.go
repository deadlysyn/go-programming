package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("go routine 0")
		wg.Done()
	}()

	go func() {
		fmt.Println("go routine 1")
		wg.Done()
	}()

	wg.Wait()
}
