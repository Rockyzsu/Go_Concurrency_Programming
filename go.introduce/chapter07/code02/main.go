package main

import (
	"fmt"

	"go.introduce/chapter07/code02/entity"
)

func main() {
	stu := entity.Student{}
	stu.New("jack", 33, "男", "CUMT")
	//stu.name
	fmt.Println(stu)
	stu.GotoSchool()
	//覆盖了Person的Walk方法
	stu.Walk()
	//继承Person的Eat方法
	stu.Eat()
}
