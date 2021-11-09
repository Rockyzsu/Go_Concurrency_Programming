package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//根据参数进行提取
	//go run .\main.go -fpath hello.txt -flen 6
	fptr := flag.String("fpath", "hello.txt", "-fpath指定文件路径读取")
	byteLen := flag.Int("flen", 6, "-flen指定读取的字节数")
	flag.Parse()
	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	//关闭文件，释放资源
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件操作失败:", err)
		}
	}()
	//文件分片读取到内存中，可以处理大文件
	r := bufio.NewReader(file)
	buffer := make([]byte, *byteLen)
	fmt.Println("文件内容为:")
	for {
		n, err := r.Read(buffer)
		if err != nil {
			//fmt.Println("文件读取完毕:", err)
			break
		}
		//fmt.Println(string(buffer)) //可能会有额外的字符
		fmt.Print(string(buffer[:n]))
	}

}
