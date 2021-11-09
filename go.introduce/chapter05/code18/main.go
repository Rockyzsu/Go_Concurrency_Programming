package main

import (
	"fmt"
)

func sum(slice []int) int {
	length := len(slice)
	ret := 0
	for i := 0; i < length; i++ {
		ret += slice[i]
	}
	return ret
}
func main() {
	var slice []int
	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}
	ret := sum(slice)
	//55
	fmt.Println(ret)
}
