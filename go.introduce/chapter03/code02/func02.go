package main

import "fmt"

func sum(a int, b int) func(int, int) int {
	return func(c int, d int) int {
		return a + b + c + d
	}
}

func main() {
	c := sum(2, 3)(4, 5)
	//sum(2, 3)=14
	fmt.Printf("sum(2, 3)(4, 5)=%d \n", c)
}
