//原子计算器2.0迭代版
//基于函数实现基本计算功能
package main

import (
	"fmt"

	"go.introduce/chapter03/code12/calc"
)

func main() {
	//质子,中子,电子
	p, n, e := 0, 0, 0
	//Scanln遇到换行时才停止扫描
	fmt.Scanln(&p, &n, &e)
	//函数多返回值
	Z, A, c := calc.AtomCalc(p, n, e)
	fmt.Printf("当质子p=%d、中子n=%d和电子e=%d时 \n", p, n, e)
	fmt.Printf("原子序数Z=%d、原子质量A=%d和电荷c=%d \n", Z, A, c)
}
