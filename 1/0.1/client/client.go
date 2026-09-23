package main

import (
	"flag"
	"fmt"
	"net"
	"io"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端
	Client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error = ", err)
		return nil
	}

	Client.conn = conn

	return Client //返回客户端
}

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址")
	flag.IntVar(&serverPort,"port",8888,"设置服务器端口")
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>输入用户名")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_,err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error=",err)
		return false
	}

	return true
}

func (client *Client) DealResponse() {
	io.Copy(os.Stdout,client.conn)//一旦client.conn有数据，就直接copy到stdout的标准输出上，永久监听
}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true{}
		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择...")
			break
		case 2:
			fmt.Println("私聊模式选择...")
			break
		case 3:
			client.UpdateName()
			break
		}
	}
}
func main() {
	flag.Parse()//解析命令行参数

	client := NewClient(serverIp,serverPort)
	if client == nil {
		fmt.Printf(">>>>> 服务器连接失败...")
		return
	}

	fmt.Println(">>>>> 服务器连接成功...")

	go client.DealResponse()
	
	//启动客户端业务
	client.Run()
}
//菜单
func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag < 4 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入合理范围内的数字<<<<")
		return false
	}
}