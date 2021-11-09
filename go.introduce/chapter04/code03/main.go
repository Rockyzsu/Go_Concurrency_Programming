//swith示例
package main

import (
	"fmt"
)

//swith获取性别
func getSex(code int) string {
	switch code {
	case 1:
		return "男"
	case 0:
		return "女"
	default:
		return "未知"
	}
}
func main() {
	code := 1
	ret := getSex(code)
	fmt.Printf("%d->%s\n", code, ret)
	code = 0
	ret = getSex(code)
	fmt.Printf("%d->%s\n", code, ret)
}
