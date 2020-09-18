package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// fixng race condition with atomic...
// (don't do this without good reason!)

func main() {
	var wg sync.WaitGroup
	var count int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
