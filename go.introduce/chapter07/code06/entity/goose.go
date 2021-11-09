package entity

import (
	"fmt"
)

type Goose struct {
	Color string
}

func (this *Goose) Sleep() {
	fmt.Println("Goose Sleep")
}
func (this *Goose) Eat() {
	fmt.Println("Goose Eat")
}
func (this *Goose) SingGua() {
	fmt.Println("Goose SingGua")
}
func (this *Goose) Type() string {
	return "Goose,Like Duck"
}
