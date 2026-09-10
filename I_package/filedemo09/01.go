package main
import (
	"fmt"
	"os"
)
func main(){
	file := "E:/text.txt"
	content,err := os.ReadFile(file)
	if err != nil {
		fmt.Println("file read error=",err)
	}
	fmt.Printf("%v",string(content))
}