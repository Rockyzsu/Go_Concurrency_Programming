package sse

import (
	"fmt"
	"log"
	"net/http"
)

//Broker 负责客户端连接注册，监听事件并广播到注册的通道上
type Broker struct {
	//Events推送到此通道上
	Notifier chan []byte
	//客户端连接
	newClients chan chan []byte
	//关闭客户端连接
	closingClients chan chan []byte
	//客户端连接注册
	clients map[chan []byte]bool
}

//监听方法
func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			//新连接时进行注册
			broker.clients[s] = true
			log.Printf("新客户端加入，共 %d 客户端", len(broker.clients))
		case s := <-broker.closingClients:
			//新连接关闭时移除注册
			delete(broker.clients, s)
			log.Printf("客户端退出. 共 %d 客户端", len(broker.clients))
		case event := <-broker.Notifier:
			//广播消息
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}
func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "不支持Stream", http.StatusInternalServerError)
		return
	}
	//设置头信息SSE
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	//每一个连接有自己的消息通道
	messageChan := make(chan []byte)
	//新连接时通知broker
	broker.newClients <- messageChan
	//连接移除时通知broker
	defer func() {
		broker.closingClients <- messageChan
	}()
	notify := rw.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	//阻塞等待
	for {
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)
		flusher.Flush()
	}
}
