//case多个表达式示例
package main

import (
	"fmt"
)

func getQuarter(m int) string {
	switch m {
	case 1, 2, 3:
		{
			return "第一季度"
		}
	case 4, 5, 6:
		{
			return "第二季度"
		}
	case 7, 8, 9:
		{
			return "第三季度"
		}
	case 10, 11, 12:
		{
			return "第四季度"
		}

	}
	return "参数必须1-12"
}
func main() {
	m := 7
	//7月是第三季度
	fmt.Printf("%d月是%s\n", m, getQuarter(m))

}
