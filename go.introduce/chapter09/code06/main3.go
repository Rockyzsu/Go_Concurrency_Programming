package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

//模拟输入流，每隔1秒产生一个数据
func inputStream(ch chan string) {
	var i = 0
	for {
		ch <- strconv.Itoa(i) + "\n"
		time.Sleep(1 * time.Second)
		i++
	}

}

//模拟输出流，并将数据流拷贝到标准输出上
func outStream(ch chan string) {
	for {
		data := <-ch
		buf := bytes.NewBufferString(data)
		//Buffer实现了io.Reader接口
		IOCopy(buf, os.Stdout)
	}
}
func IOCopy(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
}
func main() {
	//用通道实现2个goroutine间通信
	ch := make(chan string, 2)
	go inputStream(ch)
	go outStream(ch)
	var input string
	fmt.Scanln(&input)
}
