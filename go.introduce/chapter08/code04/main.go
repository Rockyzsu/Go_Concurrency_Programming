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
	//go run .\main.go -fpath hello.txt
	fptr := flag.String("fpath", "hello.txt", "-fpath指定文件路径读取")
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
	//分行读取到内存中，可以处理大文件
	fmt.Println("文件内容为:")
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	if err = sc.Err(); err != nil {
		fmt.Println("文件操作失败:", err)
	}

}
