package test

import (
	"sync"
	"sync/atomic"
)

var mlock sync.Mutex

func Counter() int32 {
	var counter int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mlock.Lock()
			counter++
			mlock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
func Counter2() int32 {
	var counter int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
