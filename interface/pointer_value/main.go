package main


import (
	"fmt"
)

// 使用值接收者实现接口 和 使用指针接收者实现接口的区别


type mover interface {
	move()
}

type sayer interface{
	say()
}

type escaper interface{ //接口嵌套
	mover
	sayer
}
type person struct{
	name string
	age int8
}




// // 使用值接收者实现接口：类型的值和指针都可以保存到接口变量中
// func (p person) move(){
// 	fmt.Printf("%s在跑~", p.name)
// }

// 使用指针接收者实现接口：只有类型指针能够保存到接口变量中
func (p *person) move(){
	fmt.Printf("%s在跑~", p.name)
}

func (p *person) say(){
	fmt.Printf("%s在叫~", p.name)
}

func main() {
	var m mover    // 定义一个mover类型的变量
	// p1 := person{ //person类型的值
	// 	name:"张三",
	// 	age:18,
	// }

	p2 := &person{ //person类型的指针
		name:"李四",
		age:20,
	}
	
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer
	s = p2
	s.say()
	fmt.Println(s)

	var e escaper
	e = p2
	e.move()
	e.say()
	fmt.Println(e)

}