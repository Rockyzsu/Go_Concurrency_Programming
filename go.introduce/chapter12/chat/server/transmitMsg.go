package server

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

//监听全局信道message,并转发数据
func TransmitMsg() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("TransmitMsg:", err)
		}
	}()
	for {
		select {
		case msg := <-message:
			strMsg := string(msg)
			ccF.Println(strMsg)
			//群发,消息中没有@符号
			if !strings.Contains(strMsg, "@") {
				if strings.Contains(strMsg, "$NUMGO$") {
					//获取NumGoroutine
					arr2 := strings.Split(strMsg, "]说>:")
					if len(arr2) == 2 {
						sender := strings.TrimLeft(arr2[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("NumGoroutine:" + strconv.Itoa(runtime.NumGoroutine()) + "\n")
								break
							}
						}
					}
				} else if strings.Contains(strMsg, "$NUMCONN$") {
					//获取CONNECTION数量
					ccF.Println("$NUMCONN$")
					arr2 := strings.Split(strMsg, "]说>:")
					if len(arr2) == 2 {
						sender := strings.TrimLeft(arr2[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("NumConn:" + strconv.Itoa(len(onlineUsers)) + "\n")
								break
							}
						}
					}

				} else {
					arr2 := strings.Split(strMsg, "]说>:")
					if len(arr2) == 2 {
						sender := strings.TrimLeft(arr2[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("群发成功\n")
							} else {
								v.perC <- append(msg, []byte("\n")...)
							}
						}
					}
				}

			} else if strings.Contains(strMsg, "@") {
				//单发 hello world@username
				arr := strings.Split(strMsg, "@")
				if len(arr) == 2 {
					arr2 := strings.Split(arr[0], "]说>:")
					//fmt.Println(arr2)
					if len(arr2) == 2 {
						sender := strings.TrimLeft(arr2[0], "[")
						for _, v := range onlineUsers {
							if v.name == strings.Trim(arr[1], " ") {
								v.perC <- []byte(arr[0] + "\n")
							} else if v.name == strings.Trim(sender, " ") {
								v.perC <- []byte("单发成功\n")
							} else {
								//v.perC <- []byte("******\n")
							}
						}
					}

				}
			} else {
				ccF.Println("未识别消息")
			}
		}
	}
}
