package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var counter = 0      //共享变量实现goroutine通讯
	var mlock sync.Mutex //互斥锁
	func() {
		for i := 0; i < 1000; i++ {
			go func() {
				mlock.Lock() //加锁
				counter++
				mlock.Unlock() //解锁
			}()
		}
	}()
	fmt.Println("执行完成,按任意键退出")
	var input string
	fmt.Scanln(&input)   //阻塞等待用户输入
	fmt.Println(counter) //1000
}
