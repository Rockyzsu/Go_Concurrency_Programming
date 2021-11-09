package property

import (
	"bufio"
	"fmt"
	"strings"

	"go.introduce/charpter08/code11/myio"
)

//Read 读取config.property,没有缓存
func Read(key string) string {
	fio, err := myio.Open("config.property")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sc := bufio.NewScanner(fio)
	line := ""
	slice := make([]string, 2)
	for sc.Scan() {
		line = sc.Text()
		//注释
		if strings.HasPrefix(line, "#") {
			continue
		}
		slice = strings.Split(line, "=")
		if len(slice) == 2 {
			if strings.Trim(slice[0], " ") == key {
				return strings.Trim(slice[1], " ")
			}
		}
	}
	err = fio.Close()
	if err != nil {
		fmt.Println(err)

	}
	return ""
}
