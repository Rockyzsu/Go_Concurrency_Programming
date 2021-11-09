//strconv示例
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//float32或者float64转string
	var f32 float32 = 3.14
	var f64 float64 = 3.14
	//FormatFloat第一个参数是float64类型
	//s1 := strconv.FormatFloat(f32, 'f', 3, 32)
	s1 := strconv.FormatFloat(float64(f32), 'f', 3, 32)
	s2 := strconv.FormatFloat(f64, 'f', 3, 64)
	fmt.Printf("|%-10s|\n", s1)
	fmt.Printf("|%-10s|\n", s2)
	//string转float32或者float64
	//返回float64
	f1, _ := strconv.ParseFloat(s1, 32)
	f2, _ := strconv.ParseFloat(s2, 64)
	fmt.Printf("|%-10.6f|\n", f1)
	fmt.Printf("|%-10.6f|\n", f2)
	//string转int
	//字符串只有整数的字符串才能转换成功，否则返回0
	//10进制,返回值int64类型
	int1, _ := strconv.ParseInt("45", 10, 32)
	int2, _ := strconv.ParseInt("45.0", 10, 64) //0
	fmt.Printf("|%-10d|\n", int1)
	fmt.Printf("|%-10d|\n", int2)
	//int转string
	//字符串只有整数的字符串才能转换成功，否则返回0
	//10进制,返回值int64类型
	s1 = strconv.Itoa(int(int1))
	//s2 = strconv.Itoa(3.14)
	//s2 = strconv.Itoa(int(3.14))
	s2 = strconv.Itoa(int(f32))
	fmt.Printf("|%-10s|\n", s1)
	fmt.Printf("|%-10s|\n", s2)
	//bool和string互转
	b := true
	s1 = strconv.FormatBool(b)
	b2, _ := strconv.ParseBool("false")
	fmt.Printf("|%-10s|\n", s1)
	fmt.Printf("|%-10t|\n", b2)
}
