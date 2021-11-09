//swith fallthrough示例
package main

import (
	"fmt"
)

func getSex(code int) string {
	ret := ""
	switch code {
	case 1:
		ret += "男"
		fallthrough //只能放于下面
	case 0:
		ret += "女"
		fallthrough
	default:
		ret += "未知"
	}
	return ret
}
func main() {
	code := 1
	ret := getSex(code)
	fmt.Printf("%d->%s\n", code, ret) //1->男女未知
	code = 0
	ret = getSex(code)
	fmt.Printf("%d->%s\n", code, ret) //0->女未知
}
