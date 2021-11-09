package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"go.introduce/chapter11/code05/sse"
)

func main() {

	broker := sse.NewSSE()
	//开启协程
	go func() {
		for {
			time.Sleep(time.Second * 2)
			data := fmt.Sprintf("==>%s", time.Now().Format("2006-01-02 15:04:05"))
			log.Println("Sending event data")
			//获取事件数据
			broker.Notifier <- []byte(data)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 1)
			if time.Now().Second()%2 == 0 {
				data := fmt.Sprintf("-->%s", time.Now().Format("2006-01-02 15:04:05"))
				log.Println("Sending event data2")
				//获取事件数据
				broker.Notifier <- []byte(data)
			}

		}
	}()
	http.HandleFunc("/", handleIndex)
	http.Handle("/sse", broker)
	log.Fatal("error:", http.ListenAndServe(":8080", nil))

}

// 返回静态html
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(writer, nil)
}
