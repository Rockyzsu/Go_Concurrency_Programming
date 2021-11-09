package server

func broadcast(userName string) {
	for _, v := range onlineUsers {
		v.AddUser <- []byte("用户[" + userName + "]已加入聊天室\n")
	}
}
