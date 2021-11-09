package mapreduce

import "strings"

//InputSplit 拆分输入为4份，进行并发计算
func InputSplit(txt string, out [4]chan<- string) {
	txt = strings.Replace(txt, "\n", ",", -1)
	txt = strings.Replace(txt, ".", ",", -1)
	txt = strings.Replace(txt, "!", ",", -1)
	arr := strings.Split(txt, ",") //,分隔
	mod := len(arr) % 4
	for i := 0; i < (4 - mod); i++ {
		arr = append(arr, "") //凑整
	}
	lenSplit := len(arr) / 4 //获取分段的个数
	for i := 0; i < 4; i++ {
		go func(ch chan<- string, lines []string) {
			for _, line := range lines {
				word := strings.Split(line, " ")
				for _, w := range word {
					ch <- w
				}
			}
			close(ch)
		}(out[i], arr[i*lenSplit:(i+1)*lenSplit])
	}
}
