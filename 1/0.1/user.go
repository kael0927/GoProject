package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

//创建用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() //获取客户端的远程地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	go user.ListenMessage() //启动监听当前User channel的goroutine

	return user
}

func (u *User) ListenMessage() { //监听当前User channel的方法 一旦有消息 就直接发送给对端客户端
	for {
		msg := <-u.C //阻塞式读取。
		u.conn.Write([]byte(msg + "\n"))
	}
}
