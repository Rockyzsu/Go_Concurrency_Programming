//变量函数体外声明示例
package main

import "fmt"

var a int

//函数体外只能声明变量，而不能对变量进行赋值
//a = 2
func main() {
	//局部变量
	var a int = 1
	//类型int，值1
	fmt.Printf("类型%T，值%d \n", a, a)
}
