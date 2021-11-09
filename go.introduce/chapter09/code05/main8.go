package main

import (
	"fmt"
	"time"
)

func generate1(ch chan<- int) {
	time.Sleep(3 * time.Second)
	ch <- 7

}
func generate2(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 8

}
func main() {
	//用通道实现2个goroutine间通信
	ch := make(chan int)
	go generate1(ch)
	ch2 := make(chan int)
	go generate2(ch2)
	for {
		select {
		case n1 := <-ch:
			fmt.Println(n1)
		case n2 := <-ch2:
			fmt.Println(n2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			goto EXIT //break无法退出for循环
		}
	}
EXIT:
	var input string
	fmt.Scanln(&input)
}
