package ydbase

import (
	"fmt"
	"strconv"
	"strings"
)

func float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', 2, 64)
}

var width = 15

//Print 格式化打印stu切片数据
func Print(stu []Student) {
	fmt.Println(strings.Repeat("-", width*5))
	fmt.Print("|Order" + strings.Repeat(" ", width-len("|Order")))
	fmt.Print("|XH" + strings.Repeat(" ", width-len("|XH")))
	fmt.Print("|Name" + strings.Repeat(" ", width-len("|Name")))
	fmt.Print("|Age" + strings.Repeat(" ", width-len("|Age")))
	fmt.Print("|Height" + strings.Repeat(" ", width-len("|Height")))
	fmt.Println()
	fmt.Println(strings.Repeat("-", width*5))
	for i, v := range stu {
		fmt.Print("|" + strconv.Itoa(i) +
			strings.Repeat(" ", width-len("|"+strconv.Itoa(i))))
		fmt.Print("|" + v.XH + strings.Repeat(" ", width-len("|"+v.XH)))
		fmt.Print("|" + v.Name +
			strings.Repeat(" ", width-len("|"+v.Name)))
		fmt.Print("|" + strconv.Itoa(v.Age) +
			strings.Repeat(" ", width-len("|"+strconv.Itoa(v.Age))))
		fmt.Print("|" + float32ToString(v.Height) +
			strings.Repeat(" ", width-len("|"+float32ToString(v.Height))))
		fmt.Println()
		fmt.Println(strings.Repeat("-", width*5))
	}
	fmt.Println("共打印", len(stu), "行")
	fmt.Println(strings.Repeat("-", width*5))
}
