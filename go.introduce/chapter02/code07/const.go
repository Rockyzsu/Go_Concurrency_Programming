//常量示例
package main

import "fmt"

func main() {
	const ver string = "1.0.0"
	const author = "jackwang"
	//一行多个常量
	const a, b = 1, 2
	//常量用作枚举
	const (
		A = "0"
		B = "1"
	)
	//常量可以不使用
	fmt.Printf("类型%T，值%s \n", ver, ver)
	fmt.Printf("类型%T，值%s \n", author, author)
}
