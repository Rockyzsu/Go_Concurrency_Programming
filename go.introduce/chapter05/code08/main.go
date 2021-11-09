package main

import (
	"fmt"
)

//匿名struct
var config struct {
	uid string
	pwd string
	url string
}

func main() {
	config.uid = "sa"
	config.pwd = "123#@"
	config.url = "127.0.0.1:3306"
	fmt.Printf("%T", config)
	//定义且要初始化
	config2 := struct {
		uid string
		pwd string
		url string
	}{"sa", "123#@", "127.0.0.1:3306"}
	fmt.Printf("%T", config2)
}
