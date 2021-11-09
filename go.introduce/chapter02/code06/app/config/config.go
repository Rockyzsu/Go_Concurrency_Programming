package config

import "fmt"

var Name = ""

func init() {
	fmt.Println("config package init() B ")
	Name = "B"
}

func init() {
	fmt.Println("config package init() A")
	Name = "A"
}
