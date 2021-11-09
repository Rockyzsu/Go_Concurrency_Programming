package main

import (
	"fmt"
)

//只写通道参数
func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

//只读通道参数
func square(ch <-chan int) {
	for n := range ch {
		fmt.Printf("square(%d)=%d\n", n, n*n)
	}
}
func main() {
	//用通道实现2个goroutine间通信
	ch := make(chan int)
	go generate(ch)
	go square(ch)
	var input string
	fmt.Scanln(&input)
}
