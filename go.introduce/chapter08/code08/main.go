package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"go.introduce/chapter08/code08/myio"
)

func main() {
	//根据参数进行提取,参数值有空格需要用""
	//go run .\main.go -fcontent "Hello Go!;你好,TypeScript!"
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()
	//创建文件
	fio, err := myio.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//string 转[]byte
	slice := strings.Split(*content, ";")
	for _, line := range slice {
		//写入行
		_, err := fmt.Fprintln(fio, line)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("写入成功")
	/////////////////////////////////
	fmt.Println("读取内容:")
	fio, err = myio.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sc := bufio.NewScanner(fio)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	err = fio.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
