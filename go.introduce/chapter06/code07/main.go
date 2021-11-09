package main

import (
	"fmt"
)

func main() {
	//指针声明未初始化,野指针
	var ptr *int
	//打印指针ptr指向的地址
	fmt.Printf("%p\n", ptr) // 0x0
	//panic: 无效的内存地址或nil指针解引用(pointer dereference)
	//*ptr = 3 //野指针不能进行*ptr取值
	fmt.Println(ptr) //<nil>
	var a int = 7
	//可以重新指向
	ptr = &a
	*ptr = 3
	fmt.Println(a) //3
}
