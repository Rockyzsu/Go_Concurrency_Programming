package main

import (
	"fmt"
)

func change(slice []int) {
	slice[1] = 99
}
func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	//合并
	slice = append(slice, slice2...)
	//[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(slice)
	change(slice)
	//[1 99 3 4 5 6 7 8 9 10]
	fmt.Println(slice)

}
