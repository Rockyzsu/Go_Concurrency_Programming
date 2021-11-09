package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("websocket at localhost:8080/echo")
	//绑定Handler
	http.Handle("/echo", websocket.Handler(Echo))
	http.HandleFunc("/", handleIndex)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

}

//Echo 为处理程序
func Echo(w *websocket.Conn) {
	var err error
	for {
		var recMsg string
		if err = websocket.Message.Receive(w, &recMsg); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端:", recMsg)
		msg := ""
		if recMsg == "猜猜年龄" {
			msg = "服务器:18岁"
		} else if recMsg == "你好" {
			msg = "服务器:你好，请问有什么可以帮你?"
		} else {
			msg = "服务器:" + recMsg
		}
		fmt.Println(msg)
		//给客户端发送消息
		if err = websocket.Message.Send(w, msg); err != nil {
			fmt.Println(err)
			break
		}
	}
}

// 返回静态html
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(writer, nil)
}
