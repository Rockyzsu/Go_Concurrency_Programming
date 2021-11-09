package main

import (
	"fmt"
)

func main() {
	//缓冲信道，容量为2
	ch := make(chan int, 2)
	ch <- 7
	ch <- 8
	ch <- 9              //超过容量2，继续写入则阻塞
	fmt.Println(len(ch)) //2
	fmt.Println(cap(ch)) //2
	n := <-ch
	fmt.Println(n)    //7
	fmt.Println(<-ch) //8
}
