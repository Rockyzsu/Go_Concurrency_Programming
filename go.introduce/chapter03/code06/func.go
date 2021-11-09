package main

import "fmt"

func main() {
	//匿名函数赋值给变量f，可以多次调用
	f := func(a int, b int) (int, int) {
		return a + b, a * b
	}
	s, m := f(2, 3)
	//f(2, 3)=5,6
	fmt.Printf("f(2, 3)=%d,%d \n", s, m)
	s, m = f(3, 7)
	//f(2, 3)=10,21
	fmt.Printf("f(2, 3)=%d,%d \n", s, m)
}
