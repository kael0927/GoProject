package main
import "fmt"
func main(){
	m := Mouse{"rog"}
	fmt.Println(m.name)
	//m.start()
	testInterface(m)
}

type USB interface{
	start()
	end()
}
type Mouse struct{
	name string
}
func (m Mouse) start(){
	fmt.Println("开始工作")
}
func (m Mouse) end(){
	fmt.Println("停止工作")
}
func testInterface(usb USB){
	usb.start()
	usb.end()
}