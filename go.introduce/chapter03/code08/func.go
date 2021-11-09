package main

import (
	"fmt"
)

//递归函数实现斐波那契数列
func fib(n int64) int64 {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	var i int64 = 8
	fmt.Printf("fib(8)= %d\n", fib(i)) //21
}
