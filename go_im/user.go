package main

import (
	"net"
	"fmt"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	Server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}
	go user.ListenMsg()
	return user
}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

func (u *User) ListenMsg() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) DoMsg(msg string) {
	if msg == "who" {
		u.Server.mapLock.Lock()
		for _,user := range u.Server.OnlineMap {
			OnlineMsg := fmt.Sprintf("[%s]%s:在线",user.Addr,user.Name)
			u.SendMsg(OnlineMsg)
		}
		u.Server.mapLock.Unlock()
	} else {
		u.Server.BroadCast(u,msg)
	}
}
