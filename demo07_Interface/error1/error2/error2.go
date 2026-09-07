package main

import (  
    "errors"
    "fmt"
    "math"
)
func circleArea(radius float64) (float64,error){
	if radius < 0 {
		return 0,errors.New("fdjdsflklfds")
	}
	return math.Pi,nil
}
func main(){
	radius := -10.0
	area,err := circleArea(radius)
	if err != nil{
		fmt.Print(err)
		return
	}
	fmt.Printf("%0.2f",area)
}