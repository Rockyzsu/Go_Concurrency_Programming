package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go showRecord()
	time.Sleep(time.Second)
	buf := make([]byte, 10000)
	runtime.Stack(buf, true)
	fmt.Println(string(buf))
}

func showRecord() {
	tiker := time.Tick(time.Second)
	for t := range tiker {
		fmt.Println(t)
	}
}
