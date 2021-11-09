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

func main() {
	//第1种方法
	var stu Student
	stu.Name = "Jack"
	fmt.Println(stu.Name, stu.Age, stu.Height)
	//第2种方法
	stu2 := Student{}
	stu2.XH = "064248"
	stu2.Name = "jack"
	stu2.Age = 32
	stu2.Height = 1.7
	stu2.Class = "3"
	fmt.Printf("%+v", stu2)
	fmt.Println()
	//第3种方法
	stu3 := Student{
		XH:     "064248",
		Name:   "jackwang",
		Age:    32,
		Height: 1.7,
		Class:  "3", //,不能少
	}
	fmt.Printf("%+v", stu3)
	fmt.Println()
	//第4种方法，new返回指针
	stu4 := new(Student)
	stu4.Name = "Jack2"
	fmt.Printf("%+v", stu4)
}
