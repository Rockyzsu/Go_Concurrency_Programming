//类型转换示例
package main

import "fmt"

func main() {
	//float32转int
	var pi float32 = 3.14
	p := int(pi)
	fmt.Println(p) //3
	//int转float32
	f := float32(p)
	fmt.Println(f / 2.5) // 1.2
	//float32转float64
	f64 := float64(f)
	fmt.Printf("%T\n", f64)   //输出类型,float64
	fmt.Printf("%.3f\n", f64) //保留3位小数,3.000
	//var b bool = true
	//bool类型不能转int
	//b1 := int(b)
	//s := "1"
	//string类型不能转int
	//p2 := int(s)
	//%-10.2f表示最小10个宽度,左对齐,保留2位小数
	//%12.3s表示最小12个宽度,右对齐,保留3位小数
	//|3.14      |       3.000|
	fmt.Printf("|%-10.2f|%12.3f|\n", pi, f64)
}
