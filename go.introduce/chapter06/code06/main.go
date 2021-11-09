package main

import "fmt"

func main() {
	var a int32 = 7
	p := &a
	pp := &p
	**pp = 9
	fmt.Println(a)     //9
	fmt.Println(*&a)   //9
	fmt.Println(*&*&a) //9
}
