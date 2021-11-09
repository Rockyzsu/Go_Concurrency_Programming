//for示例
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("==========")
	j := 10
	for j > 0 {
		fmt.Printf("--%d\n", j)
		j--

	}
	fmt.Println("==========")
	s := 10
	for {
		if s == 0 {
			goto label
		}
		s--
		fmt.Printf("%d\n", s)
	}
label:
	fmt.Printf("%d\n", j)

}
