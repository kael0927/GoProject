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
	OnlienMap  map[string]*User
	mapLock    sync.RWMutex
}

func (s *Server) NewServer(serverIp string, serverPort int) *Server {
	Server := &Server{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		OnlienMap:  make(map[string]*User),
		Message:    make(chan string),
	}
	return Server
}

func (s *Server) ListenMsger() {
	for {
		msg := <-s.Message //从广播入口取消息
		s.mapLock.Lock()   //加锁
		for _, user := range s.OnlienMap {
			user.C <- msg //投递到每个用户的收件箱
		}
		s.mapLock.Unlock() //解锁
	}
}

func (s *Server) BroadCast(user User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn)
	s.mapLock.Lock()
	s.OnlienMap[user.Name] = user
	s.mapLock.Unlock()
	go s.BroadCast(user, "已上线")
	select {}
}

func (s *Server) Start(serverIp string, serverPort int, conn net.Conn) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s %d", serverIp, serverPort))
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
