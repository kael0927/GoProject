package main
import(
	"fmt"
	"time"
)
/*
1.创建信道
2.download方法  发送信道
3.go...  等待信道返回消息
*/
var ch = make(chan string,10)
func download(url string){
	fmt.Println("start to download",url)
	time.Sleep(time.Second)
	ch <- url
}
func main(){
	for i := 0; i < 3; i++{
		go download("a.com/"+string(i+'0'))
	}
	for i := 0; i < 3; i++{
		msg := <-ch
		fmt.Println("finish",msg)
	}
	fmt.Println("Done")
}
