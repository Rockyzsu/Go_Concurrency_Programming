package entity

import "fmt"

//AoDi 结构体实现了Drive方法
type AoDi struct {
}

//Drive 实现了IDrive接口方法
func (AoDi) Drive(name string) {
	fmt.Println("Drive AoDi " + name + " Car")
}
