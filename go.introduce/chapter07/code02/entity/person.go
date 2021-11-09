package entity

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (p *Person) Walk() {
	fmt.Println("Person walk")
}
func (p *Person) Eat() {
	fmt.Println("Person Eat")
}
