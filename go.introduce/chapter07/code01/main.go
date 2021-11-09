package main

import (
	"fmt"

	"go.introduce/chapter07/code01/entity"
)

func main() {

	p := entity.Person{}
	//私有属性和方法不能访问
	//p.name = "Jack"
	//p.walk()
	p.SetName("Jack")
	p.SetAge(33)
	p.SetSex("男")
	p.Eat()              //Person Eat
	fmt.Printf("%+v", p) //{name:Jack age:33 sex:男}

}
