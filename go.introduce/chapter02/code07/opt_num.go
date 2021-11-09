//算术运算符示例
package main

import "fmt"

func main() {
	a, b := 11, 2
	//相加 +
	c := a + b
	fmt.Printf("%d+%d = %d \n", a, b, c) //11+2 = 13
	//相减 -
	c = a - b
	fmt.Printf("%d-%d = %d \n", a, b, c) //11-2 = 9
	//相乘 *
	c = a * b
	fmt.Printf("%d*%d = %d \n", a, b, c) //11*2 = 22
	//相除 / 保留整数
	c = a / b
	fmt.Printf("%d/%d = %d \n", a, b, c) //11/2 = 5
	//取余 %
	c = a % b
	fmt.Printf("%d%%%d = %d \n", a, b, c) //11%2 = 1
	//自增，unexpected ++ at end of statement
	//c = a++
	a++
	fmt.Printf("a++ = %d \n", a) //a++ = 12
	//自减
	a--
	fmt.Printf("a-- = %d \n", a) //a-- = 11
	fmt.Printf("%f \n", 2.0+3)   //5.000000
	//d := 2.0
	//fmt.Printf("%f \n", a+d) //类型不匹配
}
