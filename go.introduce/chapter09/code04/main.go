package main

import (
	"fmt"
	"time"
)

func main() {
	//go+空格+函数或者方法调用,即可创建goroutine
	go printHello("Hello Goroutine")
	time.Sleep(1 * time.Second)
}
func printHello(msg string) {
	fmt.Println(msg)
}
