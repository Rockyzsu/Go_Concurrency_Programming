package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/google/gops/agent"
)

//只写通道参数
func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
		buf := make([]byte, 10000)
		runtime.Stack(buf, true)
		fmt.Println(string(buf))
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
	if err := agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}

	//用通道实现2个goroutine间通信
	ch := make(chan int)
	go generate(ch)
	go square(ch)
	var input string
	fmt.Scanln(&input)

}
