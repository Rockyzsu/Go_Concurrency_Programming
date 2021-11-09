//error示例
package main

import (
	"errors"
	"fmt"
)

//倒数
func reciprocal(n int) (float64, error) {
	if n == 0 {
		return 0, errors.New("0不能求倒数")
	}
	//不是float64(1/n)
	return float64(1) / float64(n), nil
}
func main() {
	for i := 3; i >= 0; i-- {
		ret, err := reciprocal(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("1/%d = %.3f\n", i, ret)
		}
	}
}
