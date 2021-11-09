package main

import (
	"fmt"
)

func main() {
	//var定义
	var slice []string
	fmt.Println(len(slice))  //0
	fmt.Println(cap(slice))  //0
	fmt.Printf("%#v", slice) //[]string(nil)
	fmt.Println()
	//slice[0] = "Hello" //下标越界
	slice = make([]string, 6)
	slice[0] = "Hello"
	slice[5] = "Go"
	fmt.Printf("%#v\n", slice) //[]string{"Hello", "", "", "", "", "Go"}
	//切片字面量定义，注意不要指定长度，否则为数组
	slice2 := []int{1, 2, 3, 4, 5}
	slice1 := slice2    //切片赋值，共享底层数组
	slice1[0] = 9       //同时影响两个切片
	fmt.Println(slice2) //[9 2 3 4 5]
	//短变量定义，长度为3，容量为10
	slice3 := make([]int, 3, 10)
	fmt.Println(len(slice3)) //3
	fmt.Println(cap(slice3)) //10
	fmt.Println(slice3)      //[0 0 0]
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	fmt.Println(slice3) //[1 2 3]
}
