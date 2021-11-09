package main

import "fmt"

type callback func(int, int) int

func doAdd(a, b int, f callback) int {
	fmt.Println("f callback")
	return f(a, b)
}
func add(a, b int) int {
	fmt.Println("add running")
	return a + b
}

func main() {
	a, b := 2, 3
	fmt.Println(doAdd(a, b, add))

	fmt.Println(doAdd(a, b, func(a int, b int) int {
		fmt.Println("============")
		return a * b
	}))
}
