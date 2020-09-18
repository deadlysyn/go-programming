package main

import (
	"fmt"
	"runtime"
	"sync"
)

// creating a race condition...
// waitgroup is fine for kicking off work but if
// access to shared data is needed it's not enough...

func main() {
	var wg sync.WaitGroup
	count := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
