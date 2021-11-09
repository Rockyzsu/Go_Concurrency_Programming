package entity

import "fmt"

//BMW 结构体实现了Drive方法
type BMW struct {
}

//Drive 实现了IDrive接口方法
func (*BMW) Drive(name string) {
	fmt.Println("Drive BMW " + name + " Car")
}
