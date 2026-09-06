package main
import(
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func download(url string){
	time.Sleep(time.Millisecond*50)
	//wg.Add(1)
	fmt.Println("start to download",url)
	//模拟耗时操作
	wg.Done()
}

func main(){
	for i:=0; i < 100; i++{
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
	fmt.Println("Done!")
}