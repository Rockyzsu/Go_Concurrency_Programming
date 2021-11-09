package main

import (
	"fmt"
	"reflect"
)

type PC struct {
	Name string
}

func (this *PC) GetName() string {
	return this.Name
}
func (this PC) Sum(a, b int) int {
	return a + b
}
func main() {

	pc := PC{Name: "神州"}
	//指针类型能调用值接收者方法和指针接收者方法
	vt := reflect.ValueOf(&pc)
	vm := vt.MethodByName("GetName")
	results := vm.Call(nil)
	fmt.Println("GetName()=", results[0].String())
	vm = vt.MethodByName("Sum")
	res := vm.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println("Sum(3,5)=", res[0].Int())
	fmt.Println("=============")
	//非指针类型只能调用值接收者方法
	vt = reflect.ValueOf(pc)
	vm = vt.MethodByName("Sum")
	res = vm.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println("Sum(3,5)=", res[0].Int())
	//vm = vt.MethodByName("GetName")
	//results = vm.Call(nil) //panic错误

}
