package main

import "fmt"

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

func main() {
	var p Person = &Student{
		name: "tom",
		age:  19,
	}
	fmt.Println(p.getName())
}
