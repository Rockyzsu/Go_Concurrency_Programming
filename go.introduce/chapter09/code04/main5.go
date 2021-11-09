package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func div(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		wg.Done() //放于此，保证执行
	}()
	fmt.Printf("10/%d=%d\n", num, 10/num)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go div(i)
	}
	wg.Wait() //等待所有goroutine执行完成
}
