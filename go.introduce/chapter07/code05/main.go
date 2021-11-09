package main

import (
	"fmt"
)

type data interface{}
type Car struct {
	Color string
	Brand string
}

func main() {
	slice := make([]data, 3)
	slice[0] = 1                 // an int
	slice[1] = "Hello"           // a string
	slice[2] = Car{"Red", "BMW"} //a struct

	for i, v := range slice {
		//v.(type)不能在switch外的任何逻辑里面使用
		switch value := v.(type) {
		case int:
			{
				fmt.Printf("slice[%d] type is int[%d]\n", i, value)
			}
		case string:
			{
				fmt.Printf("slice[%d] type is string [%s]\n", i, value)
			}
		case Car:
			{
				fmt.Printf("slice[%d] type is Car [%s]\n", i, value)
			}
		default:
			{

			}
		}
	}
}
