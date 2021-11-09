package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 7
	}()
	fmt.Println(<-ch) //阻塞直到读取到值7
	//fmt.Println(<-ch) //读取超时,deadlock
	var input string
	fmt.Scanln(&input)

}
