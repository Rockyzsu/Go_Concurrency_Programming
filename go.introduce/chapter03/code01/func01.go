package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func main() {
	c := sum(2, 3)
	//sum(2, 3)=5
	fmt.Printf("sum(2, 3)=%d \n", c)
}

//逃逸分析 go build -gcflags "-l -m" .\func01.go
//println("sum(2, 3)= ", c)
