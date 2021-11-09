package main

import (
	"fmt"
	"runtime"
	"time"

	"go.introduce/chapter09/code02/onecore"
)

func main() {
	//GOMAXPROCS 设置可同时执行的最大CPU数
	//模拟单核
	runtime.GOMAXPROCS(1)
	gopher1 := onecore.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := onecore.Gopher{Name: "Gopher2", Id: 2}

	//gopher1和gopher2并发执行
	go gopher1.MakeCoffee("A")
	go gopher2.MakeCoffee("B")

	//等待，防止主协程退出，子协程也将退出
	time.Sleep(40 * time.Second) //10
	fmt.Println("============END===============")
}

// Gopher 1 Make Coffee A
// Gopher 1 Take Coffee A
// Gopher 1 Drink Coffee A
// Gopher 2 Make Coffee B
// Gopher 2 Take Coffee B
// Gopher 2 Drink Coffee B
