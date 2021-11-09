package server

//userInfo用于存储用户信息
type userInfo struct {
	name    string
	perC    chan []byte
	AddUser chan []byte //广播用户进入或退出
}
