package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//根据参数进行提取,参数值有空格需要用""
	//go run .\main.go -fcontent "Hello Go!"
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()
	//创建文件
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := f.WriteString(*content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	if n > 0 {
		fmt.Println("写入完成")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
