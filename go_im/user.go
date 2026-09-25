package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func (u *User) NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	go u.ListenMsg()
	return user
}

func (u *User) ListenMsg() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
