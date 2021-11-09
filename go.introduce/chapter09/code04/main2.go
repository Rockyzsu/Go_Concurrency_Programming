package main

import (
	"fmt"
	"time"
)

func main() {
	//go 匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("Hello Goroutine")
	time.Sleep(1 * time.Second)
}
