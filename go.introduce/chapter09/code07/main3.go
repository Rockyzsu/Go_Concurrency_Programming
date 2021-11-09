package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//go run -race main3.go
func main() {
	var counter int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1) //原子操作
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
