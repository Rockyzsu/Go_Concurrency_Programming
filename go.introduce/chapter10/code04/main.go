package main

import (
	"fmt"
	"reflect"
)

func main() {
	msg := "Hello"
	t := reflect.TypeOf(msg)
	fmt.Println(t) // string
	t = reflect.TypeOf(&msg)
	fmt.Println(t) // *string
	//如果类型的Kind不是Array, Chan, Map, Ptr和Slice,则panics
	fmt.Println(t.Elem()) //Type类型 string
	//reflect.ValueOf(msg).Elem() 错误
	v := reflect.ValueOf(&msg)
	//如果参数v的Kind不是Interface或Ptr,则panics
	fmt.Println(v.Elem()) //Value类型 Hello
	//访问未导出的结构体字段,则panics
	if v.Elem().CanInterface() {
		vs := v.Elem().Interface().(string) //string类型
		//interface转换错误: interface {} is string, not int
		//vs := v.Elem().Interface().(int)
		fmt.Println(vs) //Hello
	}
	//可设置才能修改
	if v.Elem().CanSet() {
		v.Elem().SetString("Go")
	}
	fmt.Println(msg) //Go
	//非指针，修改的是值的副本
	v = reflect.ValueOf(msg)
	//可设置才能修改
	if v.CanSet() {
		v.SetString("Go1")
	}
	fmt.Println(msg) //还是Go,不是Go1
}
