package main

import (
	"fmt"
)

func main() {
	slice := []int{3, 4, 5, 6, 7}

	for _, value := range slice {
		value = 10
		//10
		fmt.Println(value)
		//所有迭代的value的地址是一样的
		fmt.Println(&value)
	}
	fmt.Println(slice) //[3 4 5 6 7]
}
