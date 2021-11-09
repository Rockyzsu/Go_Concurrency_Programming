package main

import (
	"fmt"
)

//Address 地址结构
type Address struct {
	Province string
	City     string
	Street   string
	string   //匿名字段，字段名和类型都是string
}

//Student 学生结构
type Student struct {
	XH     string
	Name   string
	Age    int
	Height float32
	Class  string
	//匿名字段
	Address
}

func main() {
	stu := Student{
		XH:   "064248",
		Name: "Jack",
		Age:  32,
		Address: Address{
			Province: "JS",
			City:     "XZ",
			Street:   "KaiYuan",
			string:   "118", //换行必须末尾有,
		},
		Height: 1.7, Class: "3", //换行必须末尾有,
	}
	fmt.Printf("%+v", stu)
}
