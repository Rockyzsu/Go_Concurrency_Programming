package main

import (
	"fmt"

	_ "go.introduce/chapter02/code06/app"
)

func init() {
	fmt.Println("main package init()")
}
func main() {

	fmt.Println("main package main()")
}
