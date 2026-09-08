package main
import (
	"fmt"
	"os"
)
func main(){
	file,err := os.Open("D:/21")
	if err != nil {
		fmt.Println("open file error=",err)
	}
	fmt.Printf("file:%v",file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
}