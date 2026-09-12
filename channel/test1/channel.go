package main
import(
	"fmt"
	
)
func main(){

	var intChan chan int
	intChan = make(chan int,3)
	intChan <- 10
	fmt.Println(len(intChan),cap(intChan))
	intChan <- 20
	var num int
	num = <- intChan
	fmt.Println(num)
}
