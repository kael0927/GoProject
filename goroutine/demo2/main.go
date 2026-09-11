package main
import(
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int,10)
	lock sync.Mutex
	wg    sync.WaitGroup // 新增：等待组
)

func test(n int) {
	defer wg.Done() // 新增：协程结束时通知 WaitGroup 减 1

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i < 20; i++ {
		wg.Add(1) // 新增：启动协程前，通知 WaitGroup 加 1
		go test(i)
	}
	wg.Wait() // 新增：阻塞主协程，等待所有子协程执行完毕

	// 所有协程都写完了，现在加锁遍历是安全的
	
	lock.Lock()
	for i,v := range myMap {
        fmt.Printf("myMap[%v] = %v\n",i,v)
	}
	lock.Unlock()
}