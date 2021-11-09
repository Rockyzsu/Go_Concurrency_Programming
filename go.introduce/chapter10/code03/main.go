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

func (this Student) GetName() {
	fmt.Println(this.Name)
}
func (this *Student) GetAge() {
	fmt.Println(this.Age)
}
func (this Student) getId() {
	fmt.Println(this.id)
}

//printFiledMethod 打印结构体字段和方法
func printFiledMethod(o interface{}) {
	valueOf := reflect.ValueOf(o)
	//Value可以获取Type
	typeOf := valueOf.Type()
	fmt.Println(typeOf)
	if typeOf.Kind() == reflect.Struct {
		//获取结构体字段
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			if valueOf.Field(i).CanInterface() {
				//私有属性报错
				value := valueOf.Field(i).Interface()
				fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
			} else {
				fmt.Printf("%s:%v = %v\n", field.Name, field.Type, "私有字段")
			}
		}
	}
	//获取方法(不能获取私有方法)
	for i := 0; i < typeOf.NumMethod(); i++ {
		m := typeOf.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type) //GetName:func(main.Student)
	}
}
func main() {
	stu := Student{Name: "Jack", Age: 25}
	//只能获取值接收者方法
	printFiledMethod(stu)
	fmt.Println("=============")
	a := 2
	printFiledMethod(a)
	fmt.Println("=============")
	//可以获取指针接收者方法
	printFiledMethod(&stu)
}
