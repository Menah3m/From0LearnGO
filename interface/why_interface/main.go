package main

import (
	"fmt"
)

//为什么需要接口

type dog struct {}

func (d dog)say(){
	fmt.Println("汪汪汪")
}

type cat struct {}

func (c cat)say(){
	fmt.Println("喵喵喵")
}

type person struct {
	name string
}

func (p person)say(){
	fmt.Println("啊啊啊")
}

// 接口不管你是什么类型，只管你要实现什么方法
// 接口是一种抽象的类型
// 定义一个抽象的类型：只要实现了say这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

func da(arg sayer){
	arg.say()     //不管传进来的是什么，都会被打，被打就会叫
}




func main(){
	// c1 := cat{}
	// da(c1)
	// d1 := dog{}
	// da(d1)
	// p1 := person{
	// 	name:"张三",
	// }
	// da(p1)
	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name:"李四"}
	s = p2
	fmt.Println(s)

} 