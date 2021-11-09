package main

import (
	"fmt"
	"time"
)

func change(arr [1024]int) {
	//0xc00009a000
	//fmt.Printf("%p\n", &arr)
	for i, v := range arr {
		arr[i] = v * 2
	}
}
func changeByAddress(arr *[1024]int) {
	//0xc000096000
	//fmt.Printf("%p\n", arr)
	for i, v := range *arr {
		arr[i] = v * 2
	}
}
func main() {
	arr := [1024]int{}
	for i := 1; i <= 1024; i++ {
		arr[i-1] = i
	}
	//fmt.Printf("%p\n", &arr) //0xc000096000
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 10000000; i++ {
		change(arr)
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("change(arr)执行10000000次耗时：", elapsed)
	//fmt.Println(arr) //还是1..1024
	start = time.Now() // 获取当前时间
	sum = 0
	for i := 0; i < 10000000; i++ {
		changeByAddress(&arr)
		sum++
	}
	elapsed = time.Since(start)
	fmt.Println("changeByAddress(&arr)执行10000000次耗时：", elapsed)
	//fmt.Println(arr) //修改为2..2048
}
