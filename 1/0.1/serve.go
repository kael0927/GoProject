package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	user := NewUser(conn,s) //创建新用户

	user.Online()

	isLive := make(chan bool)//监听用户是否活跃的channel
	
	//接受客户端发送的消息
	go func() {
		buf := make([]byte,4098)
		n,err := conn.Read(buf)
		if n == 0 {
			user.Offline()
		}
		if err != nil && err != io.EOF{
			fmt.Println("conn read err=",err)
			return
		}

		msg := string(buf[:n-1])

		user.DoMessage(msg)

		isLive <- true//用户的任意消息，代表当前用户是一个活跃的
	}()
    //阻塞当前 Handler，保持连接不关闭
	for{
		select{
		case <- isLive:
			//当前用户是活跃的，应该重置定时器
			//不做任何事情，为了激活select，更新下面的定时器
		case <- time.After(time.Second * 10):
			//已经超时
			//将当前的User强制关闭
			user.SendMsg("你已被踢出")
			close(user.C)//销毁用到的资源
			conn.Close()//关闭连接
			return//退出当前Handler
		}
	}
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
