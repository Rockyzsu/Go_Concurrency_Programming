package main

import (
	"errors"
	"fmt"
	"reflect"
)

func print(msg string) {
	fmt.Println(msg)
}
func sum(a, b int) int {
	return a + b
}
func funcCall(m map[string]interface{}, fn string, ps ...interface{}) ([]reflect.Value,error) {
	fv := reflect.ValueOf(m[fn])
	if len(ps) != fv.Type().NumIn() {
		err := errors.New("参数个数错误")
		return nil,err
	}
	in := make([]reflect.Value, len(ps))
	for k, p := range ps {
		in[k] = reflect.ValueOf(p)
	}
	result := fv.Call(in)
	return result,nil
}

func main() {
	//注册函数
	funcs := map[string]interface{}{
		"print": print, 
		"sum": sum,
	}
	//动态调用
	funcCall(funcs, "print", "hello world")//hello world
	v, _ := funcCall(funcs, "sum", 1, 2)
	fmt.Println(v[0].Int()) //3
}
