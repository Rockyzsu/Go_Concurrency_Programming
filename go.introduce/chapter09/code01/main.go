package main

import (
	"fmt"
	"runtime"
	"time"

	"go.introduce/chapter09/code01/onecore"
)

func main() {
	//GOMAXPROCS 设置可同时执行的最大CPU数
	//模拟单核
	runtime.GOMAXPROCS(1)
	gopher1 := onecore.Gopher{Name: "Gopher1", Id: 1}
	gopher2 := onecore.Gopher{Name: "Gopher2", Id: 2}

	//非并发
	// gopher1.MakeCoffee("A")
	// gopher2.MakeCoffee("B")

	//gopher1和gopher2并发执行
	go gopher1.MakeCoffee("A")
	go gopher2.MakeCoffee("B")

	//等待，防止主协程退出，子协程也将退出
	time.Sleep(20 * time.Second)
	fmt.Println("============END===============")
}

// Gopher  1 Make Coffee A
// Gopher  2 Take Coffee A
// Gopher  2 Drink Coffee A
// Gopher  1 Has No  Coffee to Take
// Gopher  1 Has No Coffee to  Drink
