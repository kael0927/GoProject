package main
import(
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func download(url string){
	fmt.Println("start to download",url)
	time.Sleep(time.Millisecond * 500)//模拟耗时操作
	wg.Done()
}

func main(){
	for i:=0; i < 3; i++{
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")
}