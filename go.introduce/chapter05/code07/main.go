package main

import (
	"fmt"
)

//Address 地址结构
type Address struct {
	Province string
	City     string
	Street   string
	Other    string
}

//Student 学生结构
type Student struct {
	XH     string
	Name   string
	Age    int
	Height float32
	Class  string
	//struct嵌套
	Address Address
}

func main() {
	stu := Student{}
	stu.XH = "064248"
	stu.Name = "Jack"
	stu.Address.Province = "JS"
	stu.Address.City = "XZ"
	stu.Address.Street = "KaiYuan"
	fmt.Printf("%+v", stu)

}
