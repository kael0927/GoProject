# Golang
---
> 2026.9.5
## 主要内容
 - 结构和接口
 - 并发
## 反思
 - interface：
    只关心方法签名，不关心字段
    func (stu *Student) getName() string 指明隶属关系 把getName函数绑定到Student这个类型上
 - var p Person = &Student{...} 
    将具体类型 *Student 的实例，赋值给 Person 接口类型的变量 p
 - Goroutine：一个虚拟线程
   Channel：ch <- 塞数据，<-ch 取数据。当管道没有数据时，<-ch 会阻塞等待
---
> 2026.9.4
## 主要内容
 - 基础语法
 - := （可省略类型名）
 - []rune：转换为rune数组 每个字符都用int32表示
 - Map：键值对
 ## 反思
 - 切片长度可随时扩展 append
 