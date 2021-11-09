package main

import (
	"fmt"
	"reflect"
)

type IDuck interface {
	SingGua()
}
type Goose struct {
}

//实现IDuck接口
func (this Goose) SingGua() {
	fmt.Println("Goose Sing Gua")
}
func main() {

	intf := new(IDuck) //*IDuck
	intfType := reflect.TypeOf(intf).Elem()
	goose := Goose{}
	srcType := reflect.TypeOf(&goose)
	//判断是否实现IDuck接口
	if srcType.Implements(intfType) {
		//*main.Goose实现了main.IDuck接口
		fmt.Printf("%v实现了%v接口", srcType, intfType)
	} else {
		fmt.Printf("%v没有实现%v接口", srcType, intfType)
	}
	fmt.Println("\n==============")
	srcType = reflect.TypeOf(goose)
	if srcType.Implements(intfType) {
		//main.Goose实现了main.IDuck接口
		fmt.Printf("%v实现了%v接口", srcType, intfType)
	} else {
		fmt.Printf("%v没有实现%v接口", srcType, intfType)
	}
}
