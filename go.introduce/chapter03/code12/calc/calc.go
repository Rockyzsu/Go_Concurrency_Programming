package calc

import "fmt"

const title = "原子计算器2.0，输入质子、中子和电子的值，得到原子序数、原子质量和电荷"
const useage = "用法:空格分隔3个数值，换行进行计算。例如 16 16 18"

//自动调用
func init() {
	fmt.Println(title)
	fmt.Println(useage)
	fmt.Println("================")
}

//AtomCalc 计算函数
//
//输入p, n, e返回Z, A, c
func AtomCalc(p, n, e int) (int, int, int) {
	//原子序数（Z）
	Z := p
	//原子质量
	A := p + n
	//电荷
	c := p - e
	//多返回值
	return Z, A, c

}
