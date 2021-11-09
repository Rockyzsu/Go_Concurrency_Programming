package main

import (
	"fmt"
	"time"
)

func main() {
	//匿名函数赋值给变量
	printHello := func(msg string) {
		fmt.Println(msg)
	}
	go printHello("Hello Goroutine")
	time.Sleep(1 * time.Second)
}
