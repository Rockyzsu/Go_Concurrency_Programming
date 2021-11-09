package main

import (
	"fmt"
	"strconv"
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

	arr := [10]Student{}
	for i := 0; i < len(arr); i++ {
		arr[i] = Student{
			XH:    "XH" + strconv.Itoa(i),
			Name:  "Jack" + strconv.Itoa(i),
			Age:   10 + i,
			Class: strconv.Itoa(i),
		}
	}
	//Jack1
	fmt.Println(arr[1].Name)
	//12
	fmt.Println(arr[2].Age)
}
