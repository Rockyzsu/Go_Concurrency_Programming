package mapreduce

import "sync"

//Master 负责调度，将4个输入任务，分配到2个ReduceWorker上
func Master(in []<-chan map[string]int, out [2]chan<- map[string]int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	i := 0
	for _, ch := range in {
		go func(c <-chan map[string]int) {
			for m := range c {
				if i%2 == 0 {
					out[0] <- m
				} else {
					out[1] <- m
				}
			}
			wg.Done()
			i++
		}(ch)

	}
	go func() {
		wg.Wait()
		close(out[0])
		close(out[1])
	}()
}
