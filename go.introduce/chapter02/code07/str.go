//字符串示例
package main

import "fmt"

func main() {
	var msg = "I'm  here \n ========"
	//字符串是只读的
	//msg[0] = "i"
	var msg2 = `"Hello \n Go
  	Hello VSCode"
  	`
	var ch = 'a' //97
	//变量名建议不用下划线_连接，即msg_copy
	var msgCopy string //字符串零值"",不是nil
	fmt.Printf("类型%T,值%#v\n", msg, msg)
	fmt.Printf("类型%T,值%v\n", msg2, msg2)
	fmt.Printf("类型%T,值%c\n", ch, ch)
	//不能获取msg[0]的地址
	//fmt.Println(&msg[0])
	fmt.Printf("类型%T,值%#v\n", msgCopy, msgCopy)
}
