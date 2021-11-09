//break和continue语句示例
package main

import "fmt"

func main() {
	m := 10
	for {
		if m == 8 {
			break
		}
		fmt.Printf("%d\n", m)
		m--
	}
	fmt.Println("=======")
	fmt.Println("m=", m)
	for {
		m--
		if m == 7 || m == 5 {
			continue
		}
		if m <= 3 {
			break
		}
		fmt.Printf("%d\n", m)

	}

}
