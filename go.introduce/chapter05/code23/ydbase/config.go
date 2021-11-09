package ydbase

var config struct {
	uname   string
	pwd     string
	isLogin bool
}

func init() {
	config.uname = "root"
	config.pwd = "123"
	config.isLogin = false
}
