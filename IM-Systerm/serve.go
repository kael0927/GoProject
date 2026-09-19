package main
import (
	"fmt"
	"net"
)

func (this *Server) Handler(conn net.Conn){
	fmt.Println("连接成功")
}
type Server struct {
	Ip string
	Port int
}
//创建server的接口(构造函数)
func (server *Server) NewServer(ip string,port int) *Server {
	server = &Server{
		Ip:ip,
		Port:port,
	}
	return server
}
//启动服务器的接口
func (this *Server) start(){
	//socket listen

	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	}
	//close listen socket
	defer listener.Close()
	for{
		//accept
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=",err)
			continue
			
		}
		//do handler
		go Handler(conn)
	}
}