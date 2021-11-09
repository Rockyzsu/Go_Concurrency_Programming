//if判断示例
package main

import (
	"fmt"
)

//审批流程
func approve(days int) string {
	if days > 3 {
		return "总经理审批"
	} else { //else不能换行
		return "HR审批"
	}
}
func main() {
	days := 3
	ret := approve(days)
	fmt.Printf("请假%d天，需要%s\n", days, ret)
	days = 5
	ret = approve(days)
	fmt.Printf("请假%d天，需要%s\n", days, ret)
}
