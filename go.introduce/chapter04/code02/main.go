//if判断中的数据作用域示例
package main

import (
	"fmt"
)

//审批流程
func approve(days int) string {
	ret := ""
	if days > 3 {
		//if括号范围内的局部变量
		ret2 := "========="
		ret = "总经理审批"
		fmt.Println(ret2)
	} else if days <= 3 && days > 1 {
		ret = "HR审批"
	} else { //else不能换行
		ret = "自动审批"
	}
	//fmt.Println(ret2)
	return ret
}
func main() {
	days := 3
	ret := approve(days)
	fmt.Printf("请假%d天，需要%s\n", days, ret)
	days = 5
	ret = approve(days)
	fmt.Printf("请假%d天，需要%s\n", days, ret)
}
