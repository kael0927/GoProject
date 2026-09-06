package main

import (
	"fmt"
	//"go/types"
)
func main(){
	type IntNew = int
	type Gone int
	var a1 IntNew
	var a2 Gone
	fmt.Printf("%T %T",a1,a2)
}