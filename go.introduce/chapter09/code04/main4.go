package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	var counter = 0

	func() {
		for i := 0; i < 1000; i++ {
			go func() {
				counter++
			}()
		}
	}()
	fmt.Println("执行完成,按任意键退出")
	var input string
	fmt.Scanln(&input)   //阻塞等待用户输入
	fmt.Println(counter) //一般不是1000，每次不同
}
