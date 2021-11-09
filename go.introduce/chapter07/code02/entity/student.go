package entity

import (
	"fmt"
)

//继承通过组合来实现的
type Student struct {
	//小写无法继承属性和方法
	//person Person
	Person //组合Person
	school string
}

func (this *Student) GotoSchool() {
	fmt.Println(this.name, "go to ", this.school)
}
func (this *Student) Walk() {
	fmt.Println(this.name, " Walk")
}
func (this *Student) New(name string, age int, sex string, school string) {
	//继承的属性
	this.name = name
	this.age = age
	this.sex = sex
	//自己的属性
	this.school = school

}
