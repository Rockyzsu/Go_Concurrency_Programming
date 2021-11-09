package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 7
		close(ch) //关闭后,其他goroutine可继续从ch中读取
	}()
	fmt.Println(<-ch) //阻塞直到读取到值7
	fmt.Println(<-ch) //可读取到值0
	var input string
	fmt.Scanln(&input)
}
