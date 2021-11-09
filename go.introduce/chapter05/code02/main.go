package main

import (
	"fmt"
)

type mystr = string

func main() {
	var myname mystr = "Jack"
	myname = "Smith"
	//类型仍然是string
	fmt.Printf("%T \n", myname)
	myname2 := "Smith"
	fmt.Printf("%t \n", myname2 == myname) //true
	//Go内置的两个类型别名
	var a byte
	fmt.Printf("%T \n", a) //uint8
	var c rune
	fmt.Printf("%T \n", c) //int32
}
