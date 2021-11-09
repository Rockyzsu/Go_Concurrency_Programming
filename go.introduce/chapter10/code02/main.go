package main

import (
	"fmt"
	"reflect"
)

func printValue(obj interface{}) {
	v := reflect.ValueOf(obj)
	k := v.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("int value is:%d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("float32 value is:%f\n", v.Float())
	case reflect.Float64:
		fmt.Printf("float64 value is:%f\n", v.Float())
	case reflect.Bool:
		fmt.Printf("bool value is:%t\n", v.Bool())
	case reflect.String:
		fmt.Printf("string value is:%s\n", v.String())
	default:
		fmt.Printf("unknown type:%v", v)
	}
}

func main() {
	var o interface{}
	o = 3.1415
	//interface value -> reflect object
	fmt.Println(reflect.TypeOf(o))  //float64
	fmt.Println(reflect.ValueOf(o)) //3.1415
	//reflect object -> interface value
	iv := reflect.ValueOf(o).Interface()
	fmt.Printf("%+v", iv) //3.1415float64
	//interface value -> reflect object
	refObjectV := reflect.ValueOf(iv)
	fmt.Println(refObjectV.Kind())  //float64
	fmt.Println(refObjectV.Float()) //float64
	fmt.Println("================")
	printValue(o)
	o = "hello go"
	printValue(o)
	o = true
	printValue(o)
	o = 7
	printValue(o)

}
