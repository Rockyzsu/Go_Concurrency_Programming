package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

//goroutine泄漏，只接收
func consumer(ch chan int) {
	fmt.Println("执行...")
	data := <-ch
	fmt.Println(data)
}
func checkGLeak() {
	fmt.Println("Check Goroutine Leak...")
	slice := make([]int, 0)
	for {
		slice = append(slice, runtime.NumGoroutine())
		time.Sleep(1 * time.Second)

		if len(slice) >= 5 {
			n := slice[len(slice)-1] - slice[0]
			f, _ := strconv.ParseFloat(strconv.Itoa(n), 32)
			f2, _ := strconv.ParseFloat(strconv.Itoa(len(slice)), 32)
			//粗略代替直线斜率
			if f/f2 > 0.5 {
				fmt.Println("检测到Goroutine泄漏")
				os.Exit(100)
			}
		}
	}
}
func main() {
	ch := make(chan int)
	go checkGLeak()

	for {
		time.Sleep(1 * time.Second)
		go consumer(ch)
	}
}
