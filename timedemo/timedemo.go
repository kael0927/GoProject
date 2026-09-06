
package main
import (
	"fmt"
	"time"
)
func main(){
	//当前时间
	t1 := time.Now()
	fmt.Printf("%T",t1)
	fmt.Println(t1)
	//获取指定时间
	t2 := time.Date(2006,9,27,11,43,5,0,time.Local)
	fmt.Println(t2)

	s3 := "1999年01月23日"
	t3,err := time.Parse("2006年01月02日",s3)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(t3)
	fmt.Printf("%T",t3)
}