package server

import (
	"net"
)

func connWrite(conn net.Conn, user userInfo) {

	for {
		select {
		case msg1 := <-user.AddUser:
			_, _ = conn.Write(msg1)
		case msg2 := <-user.perC:
			_, _ = conn.Write(msg2)
		}
	}
}
