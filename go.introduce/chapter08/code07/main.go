package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//根据参数进行提取,参数值有空格需要用""
	//go run .\main.go -fcontent "Hello Go!;你好,TypeScript!"
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()
	//创建文件
	f, err := os.Create("hello.txt")
	//f, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	//string 转[]byte
	slice := strings.Split(*content, ";")
	for _, line := range slice {
		//写入行
		_, err := fmt.Fprintln(f, line)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("写入成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
