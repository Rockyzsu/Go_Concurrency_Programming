//数值示例
package main

import "fmt"

func main() {
	//float
	var a = 10.2
	//float
	var b = 9.0
	//int
	//var b1 = 9
	//float64和int类型不匹配
	//var c = a * b1
	var c = a * b
	var d int
	var f float32
	fmt.Printf("类型%T,值%#v\n", a, a)
	fmt.Printf("类型%T,值%#v\n", c, c)
	fmt.Printf("类型%T,值%#v\n", b, b)
	fmt.Printf("类型%T,值%#v\n", d, d)
	fmt.Printf("类型%T,值%#v\n", f, f)
}
