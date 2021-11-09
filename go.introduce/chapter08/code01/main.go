package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//文件读取到内存中，一般为小文件
	content, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("文件读取失败:", err)
		return
	} else {
		//[]byte转string
		fmt.Println("文件内容为:", string(content))
	}

}
