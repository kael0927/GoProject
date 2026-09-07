package main
import "fmt"
func main(){
	var s interface{} = new(Student)
	a := s.(*Student)
	fmt.Println(a)
	a,ok := s.(*Student)
	if ok {
		fmt.Println("1")
	}
}
type Student struct{}