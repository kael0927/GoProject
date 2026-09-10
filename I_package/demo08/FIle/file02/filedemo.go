package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
func main(){
	file,err := os.Open("E:/text.txt")
	if err != nil {
		fmt.Println("open file error=",err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}
	fmt.Println("over")
}