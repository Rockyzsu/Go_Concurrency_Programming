package main

import (
	"fmt"
)

func sum(arr [100]int) int {
	length := len(arr)
	sum := 0
	for i := 0; i < length; i++ {
		sum += arr[i]
	}
	return sum
}
func main() {
	//length := 100
	//var arr [length]int 必须是常量
	var arr [100]int
	length := len(arr)
	for i := 1; i <= length; i++ {
		arr[i-1] = i
	}
	sum := sum(arr) //5050
	fmt.Println(sum)

}
