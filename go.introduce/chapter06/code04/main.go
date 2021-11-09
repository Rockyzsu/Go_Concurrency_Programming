package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int32 = 8
	var f int64 = 20
	ptr := &a //*int32
	fmt.Println(ptr)
	ptr = (*int32)(unsafe.Pointer(&f)) //*int64 -> *int32
	fmt.Println(ptr)
	*ptr = 10
	fmt.Println(a) //8
	fmt.Println(f) //10
}
