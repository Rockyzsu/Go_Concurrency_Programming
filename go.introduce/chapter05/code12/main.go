package main

import (
	"fmt"
)

//printArr 打印[2][3]int数组
func printArr(arr [2][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
func main() {
	var twoArr [2][3]int
	twoArr[0][0] = 1
	twoArr[1][1] = 2
	twoArr[1][2] = 3
	printArr(twoArr)
	//定义并初始化
	arr2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //,不能省略
	}
	fmt.Println(arr2[1][2]) //6
}
