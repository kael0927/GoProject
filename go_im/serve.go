package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	ServerIp   string
	ServerPort int
	Message    chan string
	OnlineMap  map[string]*User
	mapLock    sync.RWMutex
}

func NewServer(serverIp string, serverPort int) *Server {
	Server := &Server{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		OnlineMap:  make(map[string]*User),
		Message:    make(chan string),
	}
	return Server
}

func (s *Server) ListenMsger() {
	for {
		msg := <-s.Message //从广播入口取消息
		s.mapLock.Lock()   //加锁
		for _, user := range s.OnlineMap {
			user.C <- msg //投递到每个用户的收件箱
		}
		s.mapLock.Unlock() //解锁
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	s.Message <- sendMsg
	
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn,s)
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()
	s.BroadCast(user, "已上线")
	for {
		buf := make([]byte,4096)
		n,err := conn.Read(buf)
		if err != nil {
			s.mapLock.Lock()
			delete(s.OnlineMap,user.Name)
			s.mapLock.Unlock()
			s.BroadCast(user,"已下线")
			return		//结束Handler协程
		} else {
			msg := string(buf[:n-1])
			user.DoMsg(msg)
		}
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ServerIp, s.ServerPort))
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	fmt.Println("创建监听器成功 ")
	defer listener.Close()
	go s.ListenMsger()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.Listen err = ", err)
			continue
		}
		go s.Handler(conn)
	}
}
