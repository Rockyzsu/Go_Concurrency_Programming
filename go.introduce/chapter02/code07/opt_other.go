//特殊运算符示例
package main

import "fmt"

func main() {
	a, b := 2, "Hello"
	//&返回变量存储地址
	var p = &a
	var p1 = &b
	fmt.Printf("&a = %x \n", p)  //&a = c000012090
	fmt.Printf("&b = %x \n", p1) //&b = c0000361f0
	//*返回变量存储地址的值
	c := *p
	c2 := *p1
	fmt.Printf("*p = %d \n", c)   //*p = 2
	fmt.Printf("*p1 = %s \n", c2) //*p1 = Hello
}
