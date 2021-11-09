package main

import (
	"fmt"
)

//闭包实现的累加器adder函数
func adder(i int) func(int) int {
	ret := i
	return func(n int) int {
		ret += n
		return ret
	}
}
func main() {
	fc := adder(2)
	// 2 + 1+2+3 -> 8
	fc(1)
	fc(2)
	fc(3)
	//13 = 8 + 5
	fmt.Println(fc(5))
	//20 = 13 + 7
	fmt.Println(fc(7))
}
