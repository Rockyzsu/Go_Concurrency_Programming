package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = 0
	var wg sync.WaitGroup
	//20并发量小，不容易出现问题
	//当并发量大时(如20000)可能就会出现问题
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
