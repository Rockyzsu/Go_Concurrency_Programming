package main

import "fmt"

//全局变量
var cname = "jack"
var age = 31

func printName() {
	//局部变量
	var cname = "smith"
	fmt.Printf("printName cname->%s\n", cname)
	age++                                  //单独一行
	fmt.Printf("printName age->%d\n", age) //32
	//fmt.Printf("printName age->%s\n", age++) //error
}
func main() {
	printName()
	fmt.Printf("main cname->%s\n", cname)
	fmt.Printf("main age->%d\n", age) //32
}
