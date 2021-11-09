package main

import (
	"fmt"
)

//Student 学生结构
type Student struct {
	XH     string
	Name   string
	Age    int
	Height float32
	Class  string
}

//CreateNew 给Student进行初始化的方法
func (stu Student) CreateNew(xh, name, class string, age int, height float32) Student {
	stu.XH = xh
	stu.Name = name
	stu.Age = age
	stu.Height = height
	stu.Class = class
	return stu
}
func main() {
	stu := Student{}
	stu = stu.CreateNew("064248", "jack", "3", 32, 1.7)
	fmt.Printf("%+v", stu)
}
