package main

import (
	"fmt"
	"unsafe"

	"go.introduce/chapter06/code05/entity"
)

func main() {
	stu := new(entity.Student) //*Student
	//stu.id = 2 //私有不可修改
	fmt.Printf("%+v \n", stu) //&{name: id:0}
	//第一个name string类型
	p := (*string)(unsafe.Pointer(stu))
	*p = "jack"               //突破第一个私有字段
	fmt.Printf("%+v \n", stu) //&{name:jack id:0}
	//第二个,需要指针运算，第一个是字符串长度为16->uintptr(16)
	ptr_id := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + uintptr(16)))

	//p2 := (*int)(unsafe.Pointer(stu))
	*ptr_id = 1               //突破第二个私有字段
	fmt.Printf("%+v \n", stu) //&{name:jack id:1}
}
