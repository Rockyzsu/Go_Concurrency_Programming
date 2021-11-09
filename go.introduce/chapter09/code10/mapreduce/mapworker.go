package mapreduce

//MapWorker 统计单词的个数，并给ReduceWorker准备数据
func MapWorker(in <-chan string, out chan<- map[string]int) {
	count := map[string]int{}
	for word := range in {
		if word != "" {
			count[word] = count[word] + 1
		}
	}
	out <- count
	close(out)
}
