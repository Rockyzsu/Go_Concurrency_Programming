//逻辑运算符示例
package main

import "fmt"

func main() {
	a, b := 11, 2
	isOk := false
	//&&逻辑与运算符
	//如果两边的操作数都是 true，则条件 true，否则为 false
	fmt.Printf("%t \n", a > b && isOk) //false
	//||逻辑或运算符
	//如果两边的操作数有一个 true，则条件 true，否则为 false
	fmt.Printf("%t \n", a > b || isOk) //true
	//!逻辑非运算符
	//如果条件为 true，则逻辑非条件 false，否则为 true
	fmt.Printf("%t \n", !isOk) //true
}
