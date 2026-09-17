package main

import (
	"fmt"
	"net" //做网络socket开发时，net包含我们需要所有的方法和函数
	"io"
)

func process(conn net.Conn) {
	defer conn.Close() //关闭conn
	//循环接受客户端发送的信息
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就阻塞在此
		fmt.Println("服务器在等待客户端%s发消息\n" + conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != io.EOF {
			fmt.Println("客户端已退出")
			return 
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //1.表示网络协议 2.表示本地监听端口
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen
	//循环等待客户链接
	for {
		fmt.Println("等待链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err=", err)
		} else {
			fmt.Printf("Accept suc  conn=%v\n", conn)
		}
		//起一个协程 为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen,suc=%v",listen)
}
