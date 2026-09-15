package main
import (
	"fmt"
	"net"//做网络socket开发时，net包含我们需要所有的方法和函数
)
func main() {
	fmt.Println("服务器开始监听")
	listen,err := net.Listen("tcp","0.0.0.0:8888")//1.表示网络协议 2.表示本地监听端口
	if err != nil {
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()//延时关闭listen
	//循环等待客户链接
	for {
		fmt.Println("等待链接...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=",err)
		} else {
			fmt.Printf("Accept suc  conn=%v\n",conn)
		}
	}
	//fmt.Printf("listen,suc=%v",listen)
}