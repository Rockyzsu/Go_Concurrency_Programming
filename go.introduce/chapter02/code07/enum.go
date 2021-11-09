//常量枚举模拟示例
package main

import "fmt"

//SEX 用int8定义了一个新的性别类型
type SEX int8

const (
	//MAN SEX类型，值为1
	MAN SEX = 1
	//FEMALE SEX类型，值为0
	FEMALE SEX = 0
)

func main() {
	const sex SEX = MAN
	//类型main.SEX，值1
	fmt.Printf("类型%T，值%v \n", sex, sex)
}
