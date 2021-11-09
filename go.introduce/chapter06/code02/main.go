package main

import (
	"fmt"
)

func main() {

	var a int64 = 8
	var ptr *int64 = &a
	fmt.Printf("%v\n", ptr)
	//通过指针修改值
	*ptr = 9
	fmt.Printf("%v\n", a) //9
	var c float32 = 3.14
	//ptr := &c //*float32和*int64类型不兼容
	ptr2 := &c
	fmt.Printf("%v\n", ptr2)
	fmt.Printf("%v\n", c) //3.14
	c = 6.28
	fmt.Printf("%v\n", *ptr2) //6.28

	d := new(int)          //开辟内存,d是*int
	fmt.Printf("%v\n", *d) //0
	*d = 9
	fmt.Printf("%v\n", *d) //9
}
