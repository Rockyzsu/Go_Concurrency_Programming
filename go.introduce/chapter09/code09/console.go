package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func genTwo() {
	slice := make([]string, 1000)
	for i := 0; i < 5; i++ {
		slice = append(slice, slice...)
	}
}
func genOne(ch chan int) {
	slice := make([]string, 100000)
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		slice = append(slice, slice...)
		genTwo()
	}
	close(ch)
}
func main() {
	// go tool pprof mem.out
	// svg
	fm, err := os.Create("mem.out")
	if err != nil {
		log.Fatal(err)
	}
	defer fm.Close()
	ch := make(chan int)
	go genOne(ch)
	for c := range ch {
		fmt.Println(c)
	}
	//放到最后再执行
	pprof.WriteHeapProfile(fm)

}
