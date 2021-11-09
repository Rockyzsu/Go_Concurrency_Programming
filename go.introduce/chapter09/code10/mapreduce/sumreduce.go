package mapreduce

import (
	"sync"
)

//SumReduce 汇总输出结果
func SumReduce(in []<-chan map[string]int) map[string]int {
	var wg sync.WaitGroup
	count := map[string]int{}
	var mLock sync.Mutex
	wg.Add(len(in))
	for i := 0; i < len(in); i++ {
		go func(n int, c <-chan map[string]int) {
			for m := range c {
				for word := range m {
					mLock.Lock()
					//concurrent map writes
					count[word] = count[word] + m[word]
					mLock.Unlock()
				}

			}
			wg.Done()
		}(i, in[i])
	}
	wg.Wait()
	return count
}
