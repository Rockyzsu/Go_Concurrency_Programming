package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[:] //切片拷贝
	slice2[0] = 9
	slice2[4] = 7
	//切片共享底层数据
	fmt.Println(slice)   //[9 2 3 4 7]
	fmt.Println(slice2)  //[9 2 3 4 7]
	s3 := slice[3:]      //从下标为3切到末尾
	fmt.Println(s3)      //[4 7]
	fmt.Println(len(s3)) //长度2
	fmt.Println(cap(s3)) //容量2
	s3 = slice[:3]       // 从0切到3(不包含3)
	fmt.Println(s3)      //[9 2 3]
	fmt.Println(len(s3)) //长度3
	fmt.Println(cap(s3)) //容量5
	s3 = slice[1:4]      //从下标为1开始，长度为3(4-1)
	fmt.Println(s3)      //[2 3 4]
	fmt.Println(len(s3)) //长度3
	fmt.Println(cap(s3)) //容量4
	//从下标为1的元素开始切，切片的长度(3-1)，容量(4-1)
	s3 = slice[1:3:4]
	fmt.Println(s3)      //[2 3]
	fmt.Println(len(s3)) //长度2
	fmt.Println(cap(s3)) //容量3
}
