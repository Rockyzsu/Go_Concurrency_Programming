//哲学的三段论示例
package main

import (
	"fmt"
	"io"

	"go.introduce/chapter04/code11/syllogism"
)

func doPanic() {
	err := recover()
	if err != nil {
		fmt.Println("doPanic()捕获到异常:", err)
	}
}
func main() {
	defer doPanic()
	fmt.Println("哲学的三段论，判断是否为有效式")
	fmt.Println("用法: 格子数(1-4)  大前提模式 小前提模式 结论模式,如 1 E E E")
	gz := 0
	s, m, p := "", "", ""
	ret := false
	for {
		n, err := fmt.Scanln(&gz, &s, &m, &p)
		fmt.Println(n, err, gz, s, m, p)
		if err == io.EOF {
			break
		}
		if n > 0 {
			ret, err = syllogism.IsValid(gz, s, m, p)
			fmt.Println(ret, err)
			if err == nil {
				if ret == true {
					fmt.Printf("%s%s%s%d是有效式", s, m, p, gz)
				} else {
					fmt.Printf("%s%s%s%d不是有效式", s, m, p, gz)
				}

			}
		}
		fmt.Println("=============")
	}

}
