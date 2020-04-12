package main

import "fmt"

// 常量声明
// const pi = 3.14
// const e = 2.7

// const (
// 	pi = 3.14
// 	e = 2.7
// )

//在定义常量时，若没有指定具体值，则默认等于上面常量的值
// const (
// 	n1 = 10
// 	n2 
// 	n3
// 	n4
// )

// iota 是Go语言中的常量计数器
// iota 在const关键字出现时被重置为0
// const 中每新增一行常量声明将使得 iota 计数加一

const (
	n1 = iota
	n2 = iota
	n3 = 100
	n4 = iota
)

const n5 = iota

const (
	a, b = iota + 1 , iota + 2
	c, d
	e, f
)

func main(){
	fmt.Println(n1, n2, n3, n4, n5, a, b, c, d, e, f)

}