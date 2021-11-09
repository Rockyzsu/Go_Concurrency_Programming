package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[1:3]
	fmt.Println(slice2)      //[2 3]
	fmt.Println(len(slice2)) //长度2
	fmt.Println(cap(slice2)) //容量4
	fmt.Println(slice2[0])   //2
	fmt.Println(slice2[1])   //3
	//fmt.Println(slice2[2])   //下标越界
	//fmt.Println(slice2[3])   //下标越界
	slice2[0] = 9       //修改会影响slice和slice2
	fmt.Println(slice)  //[1 9 3 4 5]
	fmt.Println(slice2) //[9 3]
	slice[0] = 7        //只会影响slice
	fmt.Println(slice)  //[7 9 3 4 5]
	fmt.Println(slice2) //[9 3]
}
