package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func genOne(ch chan int) {
	slice := make([]string, 100000)
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		slice = append(slice, slice...)
	}
	close(ch)
}
func main() {
	// go tool trace trace.out
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer f.Close()
	ch := make(chan int)
	go genOne(ch)
	for c := range ch {
		fmt.Println(c)
	}
	defer trace.Stop()
	fmt.Println("Trace stopped")
}
