package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[1:3]
	fmt.Println(len(slice2)) //长度2
	fmt.Println(cap(slice2)) //容量4
	//slice2和slice3是两个切片
	slice3 := append(slice2, 11) //扩容,有足够容量，共享底层数组
	fmt.Println(len(slice2))     //长度2
	fmt.Println(cap(slice2))     //容量4
	slice3[0] = 77
	fmt.Println(slice)                  //[1 77 3 11 5]
	fmt.Println(slice2)                 //[77 3]
	fmt.Println(slice3)                 //[77 3 11]
	slice3 = append(slice2, 11, 12, 13) //扩容,没有足够容量，此时不共享底层数组
	fmt.Println(len(slice3))            //长度5
	fmt.Println(cap(slice3))            //容量8=4*2
	fmt.Println(slice)                  //[1 77 3 11 5]
	fmt.Println(slice3)                 //[77 3 11 12 13]
	slice3[0] = 99
	fmt.Println(slice)  //[1 77 3 11 5]
	fmt.Println(slice3) //[99 3 11 12 13]
}
