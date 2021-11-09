//原子计算器,只能进行一次计算
//知道质子、中子和电子的数目，可以计算原子序数、原子质量和电荷
package main

import (
	"fmt"
)

func main() {
	//质子
	var p int = 0
	//中子
	n := 0
	//电子
	var e = 0
	fmt.Println("原子计算器1.0，输入质子、中子和电子的值，得到原子序数、原子质量和电荷")
	fmt.Println("用法:空格分隔3个数值，换行进行计算。例如 16 16 18")
	fmt.Println("================")
	//Scanln遇到换行时才停止扫描
	fmt.Scanln(&p, &n, &e)
	//原子序数（Z）
	Z := p
	//原子质量
	A := p + n
	//电荷
	c := p - e
	fmt.Printf("当质子p=%d、中子n=%d和电子e=%d时 \n", p, n, e)
	fmt.Printf("原子序数Z=%d、原子质量A=%d和电荷c=%d \n", Z, A, c)
}
