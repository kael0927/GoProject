package main
import(
	"fmt"
)
type Cat struct {
	Name string
	Age int
}
func main(){
	allChan := make(chan interface{},4)
	cat1 := Cat{Name:"miao",Age: 19}
	cat2 := Cat{Name:"hui",Age: 2}
	allChan <- cat1
	allChan <- cat2
	allChan <- 12

	cat11 := <- allChan
	fmt.Printf("cat11 = %T\n",cat11)
	a := cat11.(Cat)
	fmt.Println(a.Name)
}
	