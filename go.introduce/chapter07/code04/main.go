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
		//类型断言
		if value, ok := v.(int); ok {
			fmt.Printf("slice[%d] type is int[%d]\n", i, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("slice[%d] type is string [%s]\n", i, value)
		} else if value, ok := v.(Car); ok {
			fmt.Printf("slice[%d] type is Car [%s]\n", i, value)
		} else {

		}
	}
}
