package main

import (
	"fmt"
)

func main() {
	var a bool = false           //1Byte
	var b uint8 = 7              //1Byte
	var c uint8 = 6              //1Byte
	var d int64 = 8              //8Byte
	var e float32 = 3.14         //4Byte
	var f int32 = 9              //4Byte
	var g bool = true            //1Byte
	var h string = "hello world" //16Byte
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n", &b)
	fmt.Printf("%v\n", &c)
	fmt.Printf("%v\n", &d)
	fmt.Printf("%v\n", &e)
	fmt.Printf("%v\n", &f)
	fmt.Printf("%v\n", &g)
	fmt.Printf("%v\n", &h)
}
