package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	ServerIp string
	ServerPort int
	Message chan string
	OnlienMap map[string]*User
	mapLock sync.RWMutex
}

func (s *Server) NewServer(serverIp string,serverPort int) *Server{
	Server := &Server{
		ServerIp:serverIp,
		ServerPort:serverPort,
		OnlienMap:make(map[string]*User),
		Message: make(chan string),
	}
	return Server
}

func (s *Server) ListenMsger(){

}

func (s *Server) Handler(conn net.Conn) {
	NewUser := NewUser(conn)
	

}

func (s *Server) Start(serverIp string,serverPort int) {
	listener,err := net.Listen("tcp",fmt.Sprintf("%s %d",serverIp,serverPort))
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	}
	fmt.Println("创建监听器成功 ")
	defer listener.Close()
	go s.ListenMsger()
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("net.Listen err = ",err)
			continue
		}
	}
	go s.Handler(conn)
}
