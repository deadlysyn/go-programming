package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}(i)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	wg.Wait()
	close(c)

	fmt.Println("done")
}
