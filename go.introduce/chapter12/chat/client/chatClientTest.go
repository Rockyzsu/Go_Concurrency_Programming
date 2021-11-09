package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// go run .\chatClientTest.go -count 2000
func main() {
	sUrl := flag.String("url", "127.0.0.1:7777", "设置连接服务器地址")
	iCount := flag.Int("count", 2000, "设置并发数")
	flag.Parse()
	fmt.Printf("正在并发[%d]连接服务器[%s]……\n", *sUrl, *iCount)
	var wg sync.WaitGroup

	for i := 0; i < *iCount; i++ {
		time.Sleep(50 * time.Millisecond)
		wg.Add(1)
		go func() {
			conn, err := net.Dial("tcp", *sUrl)
			if err != nil {
				fmt.Println("连接服务器失败:", err)
				return
			}
			defer conn.Close()
			defer wg.Done()
			fmt.Println("连接服务器成功,输入用户名:")
			var userName string = "UName" + strconv.Itoa(i)
			_, _ = conn.Write([]byte(userName))

			buf2 := make([]byte, 4096)
			n, err := conn.Read(buf2)
			if err != nil {
				fmt.Println("连接读取失败:", err)
				return
			}
			fmt.Print(string(buf2[:n]))
			//发送消息到服务器
			go func() {
				//for {
				time.Sleep(5 * time.Second)
				_, _ = conn.Write([]byte("$NUMCONN$\n\n"))
				_, _ = conn.Write([]byte("$NUMGO$\n\n"))
				//}
			}()
			//接收服务器数据
			for {
				buffer2 := make([]byte, 4096)
				n, err := conn.Read(buffer2)
				if n == 0 {
					fmt.Println("服务器已关闭当前连接")
					return
				}
				if err != nil {
					fmt.Println("连接读取失败:", err)
					return
				}
				fmt.Print(string(buffer2[:n]))
			}
		}()
	}
	wg.Wait()

}
