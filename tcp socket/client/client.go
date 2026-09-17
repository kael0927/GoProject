package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main(){
	conn,err := net.Dial("tcp","192.168.239.1:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	defer conn.Close()//保证退出时关闭连接
	fmt.Println("client conn=",conn)
	//功能1：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin)//os.Stdin代表标准输入
	//新增for循环 不断读取终端输入
	for{
		//从终端读取一行输入，并准备发送给服务器
		line,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err=",err)
			break
		}
		//去掉输入末尾的换行符，方便对比exit
		msg := strings.TrimSpace(line)
		//判断是否退出
		if msg == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		//再将line发送给服务器
		n,err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据 并退出",n)
	}
}