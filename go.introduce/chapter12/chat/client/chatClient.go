package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
)

var ccS = color.New(color.FgHiGreen, color.Bold)  //绿色
var ccF = color.New(color.FgHiRed, color.Bold)    //红色
var ccY = color.New(color.FgHiYellow, color.Bold) //红色
// go run .\chatClient.go -uname jack
// go run .\chatClient.go -url 127.0.0.1:7777 -uname zs
// go run .\chatClient.go -uname ls
func main() {
	sUrl := flag.String("url", "127.0.0.1:7777", "设置连接服务器地址")
	userName := flag.String("uname", "UserName", "设置登录用户名")
	flag.Parse()
	fmt.Printf("正在连接服务器[%s]……\n", *sUrl)
	conn, err := net.Dial("tcp", *sUrl)
	if err != nil {
		ccF.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()
	ccS.Println("连接服务器成功")
	_, _ = conn.Write([]byte(*userName)) //写入服务器
	buf2 := make([]byte, 4096)
	n, err := conn.Read(buf2)
	if err != nil {
		ccF.Println("连接读取失败:", err)
		return
	}
	ccS.Print(string(buf2[:n])) //用户[XXX]已加入聊天室
	//发送消息到服务器
	go func() {
		for {
			buffer1 := make([]byte, 4096)
			//这里使用Stdin标准输入，因scanf无法识别空格
			n, err := os.Stdin.Read(buffer1)
			if err != nil {
				ccF.Println("os.Stdin读取失败:", err)
				continue
			}
			_, _ = conn.Write(buffer1[:n])
		}
	}()
	//接收服务器数据
	for {
		buffer2 := make([]byte, 4096)
		n, err := conn.Read(buffer2)
		if n == 0 {
			ccY.Println("服务器已关闭当前连接")
			return
		}
		if err != nil {
			ccF.Println("连接读取失败:", err)
			return
		}
		fmt.Print(string(buffer2[:n]))
	}

}
