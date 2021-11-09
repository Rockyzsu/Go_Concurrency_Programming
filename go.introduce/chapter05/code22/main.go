package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3}
	//v是副本
	for _, v := range slice {
		v = v * 2
	}
	fmt.Println(slice) //[1 2 3]
	//循环可追加
	for _, v := range slice {
		slice = append(slice, v)
	}
	//[1 2 3 1 2 3]
	fmt.Println(slice)
	map1 := map[string]string{
		"name": "jack",
		"xh":   "064248",
	}
	for k := range map1 {
		if k == "xh" {
			//循环可以删除
			delete(map1, "xh")
		}
	}
	fmt.Println(map1) //map[name:jack]
}
