package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 0
		ch <- 7
		close(ch)
	}()
	for {
		if n, ok := <-ch; ok {
			fmt.Println(n)
		} else {
			break
		}
	}
	var input string
	fmt.Scanln(&input)
}
