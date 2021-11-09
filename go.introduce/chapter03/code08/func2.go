package main

import "fmt"

//fibNoRec 非递归实现斐波那契数列
func fibNoRec(n int64) int64 {
	var f1 int64 = 0
	var f2 int64 = 1
	var i int64 = 0
	for ; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f1
}
func main() {
	fmt.Printf("fibNoRec(50)= %d\n", fibNoRec(50)) //12586269025
}
