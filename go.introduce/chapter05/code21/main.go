package main

import (
	"fmt"
)

func main() {

	arr := [3]int{1, 2, 3}
	for index, v := range arr {
		fmt.Printf("arr[%d]->%d", index, v)
		fmt.Println()
	}

	slice := []int{1, 2, 3, 4, 5}
	for index, v := range slice {
		fmt.Printf("slice[%d]->%d", index, v)
		fmt.Println()
	}
	map1 := map[string]string{
		"name": "jack",
		"xh":   "064248",
	}
	for k, v := range map1 {
		fmt.Printf("map[%s]->%s", k, v)
		fmt.Println()
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum) //6
}
