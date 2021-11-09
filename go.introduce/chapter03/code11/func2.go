package main

import (
	"fmt"
)

//c int相当于定义了函数sum的局部变量c
func sum(a, b int) (c int) {
	c = a + b
	d := 2
	fmt.Println(d)
	return //c=a+b
	//return d //2
}
func main() {
	//sum(2,3)= 5
	fmt.Println("sum(2,3)=", sum(2, 3))
}
