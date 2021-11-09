package main

import (
	"fmt"
	"net/http"
	"time"
)

//只写通道参数
func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

//只读通道参数
func square(ch <-chan int) {
	for {
		n := <-ch
		fmt.Printf("square(%d)=%d\n", n, n*n)
	}
}
func main() {
	//go如何实时监测当前进程生成的goroutine协程数量？
	//http://localhost:9876/debug/pprof/goroutine
	go func() {
		fmt.Println("pprof start...")
		//可以开一个http端口来侦听
		fmt.Println(http.ListenAndServe(":9876", nil))
	}()

	//用通道实现2个goroutine间通信
	ch := make(chan int)
	go generate(ch)
	go square(ch)
	var input string
	fmt.Scanln(&input)

}
