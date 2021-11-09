package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Exit自动退出程序\n")
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	input := ""
	str := ""
	for {
		fmt.Print(">")
		input, _ = f.ReadString('\n') //\n行分隔符
		if len(input) == 1 {
			continue //空行继续输入
		}
		fmt.Sscan(input, &str) //移除换行
		if str == "exit" {
			break
		} else {
			fmt.Printf("输入%s\n", str)
		}
	}
}
