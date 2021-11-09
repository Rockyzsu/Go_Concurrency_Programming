package server

import (
	"net"
	"time"
)

//HandleConnect 处理客户端请求
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	//信道overTime处理超时
	overTime := make(chan bool)
	bufUName := make([]byte, 4096)
	n, err := conn.Read(bufUName) //读取用户名
	if err != nil {
		ccF.Println("连接读取失败:", err)
		return
	}
	userName := string(bufUName[:n])
	perC := make(chan []byte)
	perAddUser := make(chan []byte)
	//!!!未判断用户名重复问题
	user := userInfo{name: userName, perC: perC, AddUser: perAddUser}
	onlineUsers[conn.RemoteAddr().String()] = user
	//新客户端连接后广播
	go broadcast(userName)

	//监听客户端自己的信道，conn是每个客户端独有的
	go connWrite(conn, user)

	//循环读取客户端发来的消息
	go connRead(conn, overTime)

	for {
		select {
		case <-overTime:
		case <-time.After(time.Second * 300):
			_, _ = conn.Write([]byte("已被系统踢出\n"))
			thisUser := onlineUsers[conn.RemoteAddr().String()].name
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.AddUser <- []byte("用户[" + thisUser + "]已被踢出\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
	}
}
