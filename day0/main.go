package main//声明所在的包

import (
		"fmt"
		//"reflect"
	)//引入标准库/包 处理标准输入输出

func main(){//main函数是整个程序的入口，main函数所在的包名也必须为main
	//str1 := "Golang"
	//str2 := "Go语言"
	//runeArr := []rune(str2)//每个字符占用int32 方便处理中文

	//fmt.Println(reflect.TypeOf(runeArr[2]).Kind())//获得变量的类型
	//fmt.Println(runeArr[2],string(runeArr[2]))
	//fmt.Printf("%d %c\n",str2[2],str2[2])
	//fmt.Println("len(runArr):",len(runeArr))

	//数组和切片

	//var arr = [5]int{1,2,3,4,5}
	//for i := 0; i < len(arr); i++{
	//	arr[i] += 100; 
	//}
	//fmt.Println(arr)

	//slice1 := make([]float32,0)//长度为0
	//slice2 := make([]float32,3,5)//长度为3，容量为5
	//slice2 = append(slice2,2,4,6,3)

	//fmt.Println(len(slice2),cap(slice2))
	////子切片 [start end)
	//sub1 := slice2[:3]
	//fmt.Println(sub1)
	//sub2 := slice2[1:4]
	//fmt.Println(sub2)
	//sub3 := slice2[3:]
	//fmt.Println(sub3)
	////合并切片
	//combined := append(sub1,sub2...)
	//fmt.Println(combined)

	//map
	//m1 := map[string]string{
	//	"JT" : "Male",
	//	"lin" : "Female",
	//}
	//fmt.Println(m1)
	//m1["Tom"] = "Female"
	//fmt.Println(m1["JT"])

	//point
	str := "golang";
	var p *string = &str
	*p = "hello"
	fmt.Println(str)
}