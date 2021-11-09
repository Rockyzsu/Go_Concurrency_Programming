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

func change(stu Student) {
	stu.Name = "Smith"
	stu.Age = 22
	fmt.Printf("%+v", stu)
	fmt.Println()
}
func main() {
	stu2 := Student{}
	stu2.XH = "064248"
	stu2.Name = "jack"
	stu2.Age = 32
	stu2.Height = 1.7
	stu2.Class = "3"
	change(stu2)
	fmt.Printf("%+v", stu2)

}
