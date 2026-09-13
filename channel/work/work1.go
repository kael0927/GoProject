package main

import (
	"fmt"
	"sync"
)
// 定义结果结构体，绑定 n 和它的累加和	
type Result struct {
	n int
	sum int
}
var wg sync.WaitGroup// 等待 8 个协程完成
func main() {
	numChan := make(chan int,200)
	resChan := make(chan Result,200)
// 1. 启动一个协程，将 1~200 放入 numChan
go func() {
	for i := 1; i <= 200; i++ {
		numChan <- i
		//fmt.Println("writeDone",i)
	}
	close(numChan)// 写入完毕，关闭 numChan
}()

// 2. 启动 8 个协程，从 numChan 取出数，计算并存入 resChan
for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			for n := range numChan {// 遍历完自动退出（依赖 numChan 的关闭）
				sum := 0
				for i := 1; i <= n; i++{
					sum += i
				}
				resChan <- Result{n:n,sum:sum}
			}
	}()
}

// 3. 等 8 个协程全部干完活后，关闭 resChan
go func() {
	wg.Wait()
	close(resChan)// 关闭结果管道（注意：谁最后写入，谁关闭）
}()

// 4. 遍历 resChan 打印结果
// 注意：因为 8 个协程并发写入，打印顺序可能是乱序的，但数据不会错
for res := range resChan {
	fmt.Printf("Result[%v] = %v\n",res.n,res.sum)
}

}