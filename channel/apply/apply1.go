package main
import(
	"fmt"
	"time"
)

func intNum(intChan chan int) {
	for i := 1; i <= 800; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int,primeChan chan int,exisChan chan bool) {
	for{
		time.Sleep(time.Microsecond*10)//模拟耗时
		num,ok := <- intChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			if num < 2 {
				continue  // 跳过 1 和 0
			}
			
			if num % i == 0{
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}
	fmt.Println("关闭一个")
	exisChan <- true

}

func main() {
	intChan := make(chan int,100)
	primeChan := make(chan int,200)
	exisChan := make(chan bool,4)

	go intNum(intChan)

	for range 4{
		go primeNum(intChan,primeChan,exisChan)
	}
	go func() {
		for range 4{
			<- exisChan
		}
		close(primeChan)
	}()
	for i := range primeChan {
		fmt.Printf("素数：%d\n",i)
	}
	fmt.Println("over")
}