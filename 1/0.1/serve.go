package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	OnlineMap map[string]*User //在线用户列表
	mapLock sync.RWMutex //保护OnlineMap的读写锁
	Message chan string //全局广播消息通道
}

// 创建server的接口(构造函数)
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		OnlineMap:  make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn) //创建新用户

	//将用户加入在线列表 (加锁保护)
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	//广播该用户上线的消息
	s.BroadCast(user,"已上线")

	//接受客户端发送的消息
	go func() {
		buf := make([]byte,4098)
		n,err := conn.Read(buf)
		if n == 0 {
			s.BroadCast(user,"下线")
		}
		if err != nil && err != io.EOF{
			fmt.Println("conn read err=",err)
			return
		}

		msg := string(buf[:n-1])

		s.BroadCast(user,msg)

	}()

	select{} //阻塞当前 Handler，保持连接不关闭
}

func (s *Server) BroadCast(user *User,msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name +msg
	s.Message <- sendMsg
}

// 监听全局广播消息的协程 一旦有消息就发送给在线User
func (s *Server) ListenMasseger() {
	for{
		msg := <- s.Message// 从广播通道取消息
		s.mapLock.Lock()// 加锁，遍历所有在线用户，将消息塞入各自的私有通道
		for _,cli := range s.OnlineMap {
			cli.C <- msg 
		}
		s.mapLock.Unlock()
	}
}



// 启动服务器的接口
func (s *Server) Start() {
	//socket listen

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//close listen socket
	defer listener.Close()

	go s.ListenMasseger()  //启动监听Message的goroutine
	
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue

		}
		//do handler
		go s.Handler(conn)
	}
}
