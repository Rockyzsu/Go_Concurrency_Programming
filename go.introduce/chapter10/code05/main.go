package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	id   string
	Name string
	Age  int
}

func main() {
	stu := Student{Name: "Jack", Age: 25}
	//参数必须为指针地址
	v2 := reflect.ValueOf(&stu.Name)
	v2.Elem().SetString("Jack2")
	fmt.Println(stu)
	v2 = reflect.ValueOf(&stu.Age)
	v2.Elem().SetInt(33)
	//包外不能进行私有变量访问
	v2 = reflect.ValueOf(&stu.id)
	v2.Elem().SetString("001")
	fmt.Println(stu)
}
