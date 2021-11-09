package main

import "fmt"

func main() {
	//匿名函数只能调用一次
	s, m := func(a int, b int) (int, int) {
		return a + b, a * b
	}(2, 3)
	//匿名函数(2, 3)=5,6
	fmt.Printf("匿名函数(2, 3)=%d,%d \n", s, m)

}
