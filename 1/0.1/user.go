package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	server *Server
}

//创建用户的API
func NewUser(conn net.Conn,s *Server) *User {
	userAddr := conn.RemoteAddr().String() //获取客户端的远程地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
		server: s,
	}
	go user.ListenMessage() //启动监听当前User channel的goroutine

	return user
}
//用户的上线业务
func (u *User) Online() {
	//将用户加入在线列表 (加锁保护)
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	
	u.server.BroadCast(u,"已上线")//广播该用户上线的消息
}
//用户的下线业务
func (u *User) Offline() {
	//将用户从在线列表中删除 (加锁保护)
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap,u.Name)
	u.server.mapLock.Unlock()
	
	u.server.BroadCast(u,"已下线")//广播该用户下线的消息
}
//用户处理消息的业务
func (u *User) DoMessage(msg string) { 
	u.server.BroadCast(u,msg)
}
//监听当前User channel的方法 一旦有消息 就直接发送给对端客户端
func (u *User) ListenMessage() { 
	for {
		msg := <-u.C //阻塞式读取。
		u.conn.Write([]byte(msg + "\n"))
	}
}
