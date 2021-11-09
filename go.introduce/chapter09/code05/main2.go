package main

import (
	"fmt"
)

func main() {
	//用通道实现2个goroutine间通信
	ch := make(chan int) //通道 chan int类型
	ch <- 7
	//fatal error: all goroutines are asleep - deadlock!
	n := <-ch
	fmt.Println(n)
	var input string
	fmt.Scanln(&input)

}
