package main

import "fmt"

//... 表示一个变长函数参数
func sum(ns ...int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return ret
}

func main() {
	fmt.Println(sum())        // 0
	fmt.Println(sum(1))       // 1
	fmt.Println(sum(2, 3))    // 5
	fmt.Println(sum(2, 3, 4)) // 9
	// nums := []int{1, 2, 3, 4}
	// sum(nums...)
}
