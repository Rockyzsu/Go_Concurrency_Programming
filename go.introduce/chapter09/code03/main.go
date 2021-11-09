package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"go.introduce/chapter09/code03/parallel"
)

func main() {
	//模拟2核
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2) //等待2个协程运行完成
	gopher1 := parallel.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := parallel.Gopher{Name: "Gopher2", Id: 2}
	t := time.Now()
	//gopher1和gopher2并发执行
	go gopher1.MakeCoffee("A", &wg)
	go gopher2.MakeCoffee("B", &wg)
	wg.Wait()
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
	fmt.Println("============END===============")
}
