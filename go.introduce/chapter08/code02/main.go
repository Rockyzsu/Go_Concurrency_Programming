package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	//根据参数进行提取
	//go run .\main.go -fpath hello.txt
	fptr := flag.String("fpath", "hello.txt", "-fpath指定文件路径读取")
	flag.Parse()
	//文件读取到内存中，一般为小文件
	content, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("文件读取失败:", err)
		return
	} else {
		//[]byte转string
		fmt.Println("文件内容为:", string(content))
	}

}
