package main

import (
	"fmt"
	"time"

	"go.introduce/chapter09/code10/mapreduce"
)

func main() {
	//go run .\main.go -race
	t := time.Now()
	input := "Go is awesome,Go is Best,Best are Good!\nYou are Hero,Thank You."
	
	fmt.Println(mapreduce.Run(input))
	elapsed := time.Since(t)
	fmt.Println("time elapsed:", elapsed)
}
