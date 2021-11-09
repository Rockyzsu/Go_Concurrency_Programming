package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	num := 3.14
	fmt.Println("type:", reflect.TypeOf(num))   //float64
	fmt.Println("value:", reflect.ValueOf(num)) //3.14
	ti := time.Now()
	fmt.Println("type:", reflect.TypeOf(ti)) //time.Time
	fmt.Println("value:", reflect.ValueOf(ti))
	//空接口类型
	var v1 interface{}
	v1 = "hello world"
	fmt.Println("type:", reflect.TypeOf(v1))   //string
	fmt.Println("value:", reflect.ValueOf(v1)) //hello world
	v1 = true
	fmt.Println("type:", reflect.TypeOf(v1))   //bool
	fmt.Println("value:", reflect.ValueOf(v1)) //true
	v1 = 2
	fmt.Println("type:", reflect.TypeOf(v1))   //int
	fmt.Println("value:", reflect.ValueOf(v1)) //2
	v1 = make([]string, 0)
	fmt.Println("type:", reflect.TypeOf(v1))                   //[]string
	fmt.Println("type kind:", reflect.TypeOf(v1).Kind())       //slice
	fmt.Println("value:", reflect.ValueOf(v1))                 //[]
	fmt.Println("value CanSet:", reflect.ValueOf(v1).CanSet()) //false
	v1 = nil
	fmt.Println("type:", reflect.TypeOf(v1))   //<nil>
	fmt.Println("value:", reflect.ValueOf(v1)) //<invalid reflect.Value>
}
