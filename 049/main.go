package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Start goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("go routine 0")
		wg.Done()
	}()

	go func() {
		fmt.Println("go routine 1")
		wg.Done()
	}()

	fmt.Println("Middle goroutines:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End goroutines:", runtime.NumGoroutine())
}
