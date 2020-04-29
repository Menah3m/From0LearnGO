package main

import "fmt"

//指针

func main(){
	a := 10   //定义一个int类型的变量a
	b := &a   //b是一个指针变量，b的类型为 *int

	fmt.Printf("%v %p\n", a, &a) //10 0xc0000120b0
	fmt.Printf("%v %T\n", b, b) //0xc00011c068 *int
	// 变量b是一个int类型的指针（*int），它存储的是变量a的内存地址

	c := *b //根据变量b中存储的内存地址取值
	fmt.Println(c)
}