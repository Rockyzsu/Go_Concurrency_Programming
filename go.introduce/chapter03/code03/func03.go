package main

import "fmt"

//sumAndmul 返回两个值，一个是和，一个是乘
func sumAndmul(a int, b int) (int, int) {
	return a + b, a * b
}
func main() {
	s, m := sumAndmul(2, 3)
	//sumAndmul(2, 3)=5,6
	fmt.Printf("sumAndmul(2, 3)=%d,%d \n", s, m)
	s1, _ := sumAndmul(2, 3)
	//sumAndmul(2, 3)->5
	fmt.Printf("sumAndmul(2, 3)->%d\n", s1)
	//可以不赋值
	sumAndmul(2, 3)
}
