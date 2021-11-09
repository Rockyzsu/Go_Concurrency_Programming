package entity

import (
	"fmt"
)

type Duck struct {
	Color string
	Age   int
}

func (this *Duck) Sleep() {
	fmt.Println("Duck Sleep")
}
func (this *Duck) Eat() {
	fmt.Println("Duck Eat")
}
func (this *Duck) SingGua() {
	fmt.Println("Duck SingGua")
}
func (this *Duck) Type() string {
	return "Duck"
}
