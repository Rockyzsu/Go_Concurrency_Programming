//panic示例
package main

import (
	"fmt"
)

//当捕获panic异常时触发
func doPanic() {
	err := recover() //recover捕获panic异常
	if err != nil {
		//runtime error: integer divide by zero
		fmt.Println("doPanic()捕获到异常:", err)
	}
}
func main() {
	//注册捕获异常的处理函数
	defer doPanic()
	n := 0
	ret := 1 / n //抛出panic异常,不能直接写1/0，编码阶段报错
	//defer doPanic()//必须先注册
	fmt.Println("main ret =", ret) //panic异常后不会打印
}
