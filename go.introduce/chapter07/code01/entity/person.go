package entity

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (this *Person) SetName(name string) {
	this.name = name
}
func (this *Person) SetAge(age int) {
	if age > 0 && age < 200 {
		this.age = age
	}
}
func (this *Person) SetSex(sex string) {
	if sex == "男" || sex == "女" {
		this.sex = sex
	} else {
		this.sex = "未知"
	}

}
func (this *Person) GetName() string {
	return this.name
}
func (this *Person) GetAge() int {
	return this.age
}
func (this *Person) GetSex() string {
	return this.sex
}

//私有方法
func (p *Person) walk() {
	fmt.Println("Person walk")
}

//Eat 公有方法
func (p *Person) Eat() {
	fmt.Println("Person Eat")
}
