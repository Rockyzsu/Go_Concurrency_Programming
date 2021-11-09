package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i) //数据竞争
			wg.Done()
		}()
	}
	wg.Wait()
}
