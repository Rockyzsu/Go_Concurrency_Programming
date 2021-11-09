package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//goroutine泄漏,mlock未释放
func adder() {
	fmt.Println("执行...")
	mlock.Lock()
	counter++
	fmt.Println("counter=", counter)
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

var mlock sync.Mutex
var counter = 0

func main() {
	go checkGLeak()
	for {
		time.Sleep(1 * time.Second)
		go adder()
	}
}
