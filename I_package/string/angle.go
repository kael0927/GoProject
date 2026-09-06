package main
import "fmt"
func main(){
	angle := "hero never die"
	s := []byte(angle)
	for i := 5; i < 10; i++{
		s[i] = ' '
	}
	fmt.Println(string(s))
}