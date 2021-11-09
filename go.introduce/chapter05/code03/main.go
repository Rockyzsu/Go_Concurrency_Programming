//自定义类型的方法示例
package main

import (
	"fmt"
	"strconv"
)

//自定义类型str和myInt
type str string
type myInt int

//getLen str类型的方法,获取自身的长度
func (m str) getLen() int {
	return len(m)
}
func (i myInt) toStr() string {
	return strconv.Itoa(int(i))
}

//接收器类型int不合法,基础类型(basic type)不支持
// func (i int) toStr() string {
// 	return strconv.Itoa(i)
// }

func main() {
	var name str = "Jack"
	fmt.Printf("%d\n", name.getLen()) //4
	msg := "Hello," + name
	fmt.Printf("%T\n", msg)          //main.str
	fmt.Printf("%d\n", msg.getLen()) //10

	var i myInt = 30
	i2 := i + 10
	fmt.Printf("%T\n", i2)          //myInt
	fmt.Printf("%#v\n", i.toStr())  //"30"
	fmt.Printf("%#v\n", i2.toStr()) //"30"
}
