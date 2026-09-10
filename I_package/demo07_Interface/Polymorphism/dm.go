package main

import "fmt"

type Income interface {
	cal() int
	src() string
}
type FixedBilling struct {
	projectName  string
	biddedAmount int
}
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fi FixedBilling) cal() int {
	return fi.biddedAmount
}
func (fi FixedBilling) src() string {
	return fi.projectName
}
func (ti TimeAndMaterial) cal() int {
	return ti.noOfHours * ti.hourlyRate
}
func (ti TimeAndMaterial) src() string{
	return ti.projectName
}
func num (ic []Income){
	netincome := 0
	for _,income := range ic{
		fmt.Printf("Income From %s = $%d\n", income.src(), income.cal())
        netincome += income.cal()
	}
	 fmt.Printf("Net income of organisation = $%d", netincome)
}
func main(){
	project1 := FixedBilling{
		projectName:"klfdsj",
		biddedAmount: 123,
	}
	project2 := TimeAndMaterial{
		projectName:"324",
		noOfHours: 1233,
		hourlyRate: 23,
	}
	incomestream := []Income{project1,project2}
	num(incomestream)
}
