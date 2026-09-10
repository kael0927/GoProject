package main
import (
	"fmt"
)
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++{
		res += i
	}
	return res
}
func main(){
	res := addUpper(10)
	if res != 55{
		fmt.Printf("addUpper错误 返回值%v 期望值%v",res,55)
	}
}