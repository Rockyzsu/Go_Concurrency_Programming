package main

import (
	"fmt"
	"runtime"
	"time"
)

//goroutine泄漏，只发送没有接收
func consumer(ch chan int) {
	for {
		data := <-ch
		fmt.Println(data)
	}
}
func main() {
	ch := make(chan int)
	for {
		//模拟持续运行
		time.Sleep(1 * time.Second)
		go consumer(ch)
		//goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
		buf := make([]byte, 10000)
		runtime.Stack(buf, true)
		fmt.Println(string(buf))
		fmt.Println(string(buf))
	}
}
