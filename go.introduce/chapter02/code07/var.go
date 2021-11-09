//变量示例
package main

import "fmt"

func main() {
	var ver string = "1.0.0"
	var author = "jackwang"
	time := "2019-12-1"
	//一行多个变量
	var a, b int = 1, 2
	var (
		A string = "0"
		B        = "1"
	)
	fmt.Printf("类型%T，值%s \n", ver, ver)
	fmt.Printf("类型%T，值%s \n", author, author)
	fmt.Printf("类型%T，值%s \n", time, time)
	fmt.Printf("类型%T，值%d \n", a, a)
	fmt.Printf("类型%T，值%d \n", b, b)
	fmt.Printf("类型%T，值%s \n", A, A)
	fmt.Printf("类型%T，值%s \n", B, B)
}
