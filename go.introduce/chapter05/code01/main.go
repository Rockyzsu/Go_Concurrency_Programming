package main

import (
	"fmt"
)

//定义一个新的类型name
type name string

func main() {
	var myname name = "Jack"
	myname2 := myname
	//main.name
	fmt.Printf("%T \n", myname2)
	fmt.Printf("%t \n", myname2 == "Jack") //true
	//myname3 := "Jack"
	//类型name和string不匹配，无法比较
	//fmt.Printf("%t \n", myname2 == myname3) //错误
}
