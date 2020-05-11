package main

import "fmt"

// 结构体的初始化

type person struct {
	name, city string
	age		   int8
}

func main() {
	// 1. k-v对初始化（某些key可以省略，则默认为零值）
	p4 := person{
		name: "小王",
		city: "北京",
		age: 18,
}
	fmt.Printf("%#v\n", p4)

	p5 := &person{
		name: "小王",
		city: "北京",
		age: 18,
}
	fmt.Printf("%#v\n", p5)

	// 2. 值的列表进行初始化（必须要一一对应）
	p6 := person{
		"校长",
		"成都",
		35,
	}
	fmt.Printf("%#v\n", p6)
}