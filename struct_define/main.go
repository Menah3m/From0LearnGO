package main

import "fmt"

// 定义结构体

type person struct {
	name, city string
	age		   int8
}

// 结构体的实例化

func main() {
	var p1 person
	p1.name = "张三"
	p1.city = "成都"
	p1.age = 18
	fmt.Printf("p1=%#v\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.city)
	fmt.Println(p1.age)

	// 匿名结构体
	var user struct {
		name string
		married bool
	}

	user.name = "小王"
	user.married = false
	fmt.Printf("user=%#v\n", user)
}