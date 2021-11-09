package main

import (
	"fmt"
)

func print(s string) {
	fmt.Println("run ", s)
}
func main() {
	fmt.Println("====start=====")
	//defer执行顺序和调用顺序相反
	defer print("order 1")
	defer print("order 2")
	defer print("order 3")
	fmt.Println("====end=====")
}

// ====start=====
// ====end=====
// run  order 3
// run  order 2
// run  order 1
