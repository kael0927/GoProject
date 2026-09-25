package main

import (
	"net"
	"fmt"
	"strings"	
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
	} else if len(msg) >= 7 && msg[:7] == "rename|" {
		newName := strings.TrimSpace(msg[7:])
		_,ok := u.Server.OnlineMap[newName] 
		if ok  {
			u.SendMsg("当前用户名被使用")
		} else {
			u.Server.mapLock.Lock()
			delete(u.Server.OnlineMap,u.Name)
			u.Server.OnlineMap[newName] = u
			u.Server.mapLock.Unlock()
			u.Name = newName
			u.SendMsg("你已经更新用户名：" + u.Name)
		}
	} else if len(msg) >=3 && msg[:3] == "to|" {
		parts := strings.Split(msg,"|")
		user,ok := u.Server.OnlineMap[parts[1]]
		if !ok {
			u.SendMsg("用户名不存在")
		} else {
			msg := fmt.Sprintf("[私聊]%s:%s",u.Name,parts[2])
			user.C <- msg
		}
	} else {
		u.Server.BroadCast(u,msg)
	}
}
