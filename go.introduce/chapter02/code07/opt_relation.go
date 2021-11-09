//关系运算符示例
package main

import "fmt"

func main() {
	a, b := 11, 2
	//==检查两个值是否相等，如果相等返回 true 否则返回 false
	c := a == b
	fmt.Printf("%d==%d ? %t \n", a, b, c) //11==2 ? false
	//!=检查两个值是否不相等，如果不相等返回 true 否则返回 false
	c = a != b
	fmt.Printf("%d!=%d ? %t \n", a, b, c) //11!=2 ? true
	//> 检查左边值是否大于右边值，如果是返回 true 否则返回 false
	c = a > b
	fmt.Printf("%d>%d ? %t \n", a, b, c) //11>2 ? true
	//>= 检查左边值是否大于等于右边值，如果是返回 true 否则返回 false
	c = a >= b
	fmt.Printf("%d>=%d ? %t \n", a, b, c) //11>=2 ? true
	//< 检查左边值是否小于右边值，如果是返回 true 否则返回 false
	c = a < b
	fmt.Printf("%d<%d ? %t \n", a, b, c) //11<2 ? false
	//<= 检查左边值是否小于等于右边值，如果是返回 true 否则返回 false
	c = a <= b
	fmt.Printf("%d<=%d ? %t \n", a, b, c) //11<=2 ? false
	fmt.Printf("%t \n", 2.0 < 3)          //true
	//d := 2.0
	//fmt.Printf("%t \n", d < a) //类型不匹配
}
