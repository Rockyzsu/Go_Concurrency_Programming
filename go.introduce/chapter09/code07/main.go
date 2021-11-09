package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 7
	}()
	ch <- 8
	fmt.Println(<-ch)
}
