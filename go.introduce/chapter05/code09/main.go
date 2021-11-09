package main

import (
	"fmt"
)

func main() {
	//定义一个长度为3,元素类型为string的数组
	//初始化为string零值""
	var arrName [3]string
	fmt.Printf("%#v", arrName) //[3]string{"", "", ""}
	fmt.Println()
	//通过下标赋值
	arrName[0] = "Hello"
	arrName[1] = ","
	arrName[2] = "Go"
	fmt.Println(arrName) //[Hello , Go]
	//短变量定义,长度为3,元素类型为int的数组
	//初始化为int零值0
	arrInt := [3]int{}
	fmt.Println(arrInt) //[0 0 0]
	arrInt[0] = 1
	arrInt[1] = 2
	arrInt[2] = 3
	//越界报错
	//arrInt[4] = 4
	fmt.Println(arrInt) //[1 2 3]
	//短变量定义并初始化,长度为3,元素类型为int的数组
	arr2 := [3]int{7, 8, 9}
	fmt.Println(arr2) //[7 8 9]
	//可以用...,会自动根据元素个数确定数组长度
	arr3 := [...]int{7, 8, 9, 10}
	fmt.Println(arr3) //[7 8 9 10]
}
