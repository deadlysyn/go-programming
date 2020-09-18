package main

import (
	"fmt"
	"sync"
)

// fixng race condition with mutex...

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	count := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
