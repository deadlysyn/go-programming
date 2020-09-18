package main

import (
	"fmt"
	"runtime"
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
			runtime.Gosched()
			v++
			count = v
			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
