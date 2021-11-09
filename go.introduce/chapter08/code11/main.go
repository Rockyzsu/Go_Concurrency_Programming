package main

import (
	"flag"
	"fmt"

	"go.introduce/chapter08/code11/property"
)

func main() {
	//go run .\main.go -fkey db.uname
	content := flag.String("fkey", "db.uname", "-fkey指定配置文件的key")
	flag.Parse()

	fmt.Printf("key[%s]->%s", *content, property.Read(*content))

}
