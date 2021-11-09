//赋值运算符示例
package main

import "fmt"

func main() {
	//二进制数值表示a=60,b = 141
	a, b := 0b00111100, 0b10001101
	var c int
	//= 简单的赋值运算符
	c = a + b
	fmt.Printf("%b + %b = %b \n", a, b, c) //111100 + 10001101 = 11001001
	fmt.Printf("%d + %d = %d \n", a, b, c) //60 + 141 = 201
	c1 := c
	//相加后再赋值
	c += a
	fmt.Printf("%d += %d -> %d \n", c1, a, c) //201 += 60 ->  261
	c1 = c
	//相减后再赋值
	c -= a
	fmt.Printf("%d -= %d -> %d \n", c1, a, c) //201 -= 60 ->  201
	c1 = c
	//相乘后再赋值
	c *= a
	fmt.Printf("%d *= %d -> %d \n", c1, a, c) //201 *= 60 ->  12060
	c1 = c
	//相除后再赋值
	c /= a
	fmt.Printf("%d /= %d -> %d \n", c1, a, c) //201 /= 60 -> 3
	c1 = c
	//求余后再赋值
	c %= a
	fmt.Printf("%d %%= %d -> %d \n", c1, a, c) //3 %= 60 -> 3
	c1 = a
	//左移后赋值
	a <<= 3
	fmt.Printf("%b <<=3 -> %b \n", c1, a) //111100 <<=3 -> 111100000
	c1 = a
	//右移后赋值
	a >>= 3
	fmt.Printf("%b >>=3 -> %b \n", c1, a) //111100000 >>=3 -> 111100
	c1 = a
	//按位与后赋值
	a &= b
	fmt.Printf("%b &= %b -> %b \n", c1, b, a) //111100 &= 10001101 -> 1100
	c1 = a
	//按位异或后赋值
	a ^= b
	fmt.Printf("%b ^= %b -> %b \n", c1, b, a) //1100 ^= 10001101 -> 10000001
	c1 = a
	//按位或后赋值
	a |= b
	fmt.Printf("%b |= %b -> %b \n", c1, b, a) //10000001 |= 10001101 -> 10001101
}
