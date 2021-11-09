package main

import (
	"fmt"
)

func change(arr [10]int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		arr[i] = 2 * i
	}

}
func main() {

	var arr [10]int
	length := len(arr)
	for i := 1; i <= length; i++ {
		arr[i-1] = i
	}
	change(arr)
	//[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(arr)
}
