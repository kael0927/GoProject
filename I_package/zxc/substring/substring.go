package main
import (
	"fmt"
	"strings"
)
func main(){
	tracer := "这辈子，就学计算机"
	comma := strings.Index(tracer,"，")
	pos := strings.Index(tracer[comma:],"就学")
	fmt.Println(comma,pos,tracer[comma+pos:])
}