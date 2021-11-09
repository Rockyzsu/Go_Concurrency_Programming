package main

import (
	"fmt"
	"reflect"
)

//Student 字段标识Tag
type Student struct {
	id   string `json:"id" iskey:"1"`
	Name string `json:"cname" table:"t_student"`
	Age  int    `json:"age"`
}

func main() {
	stu := Student{Name: "Jack", Age: 25}
	t := reflect.TypeOf(&stu).Elem()
	m := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = t.Field(i).Tag.Get("json")
	}
	fmt.Println(m) //map[Age:age Name:cname id:id]
	if f, ok := t.FieldByName("Name"); ok {
		//t_student
		fmt.Println("Name字段的table标识值:", f.Tag.Get("table"))
	}
	if f, ok := t.FieldByName("id"); ok {
		//1
		fmt.Println("id字段的iskey标识值:", f.Tag.Get("iskey"))
	}
}
