package main

import "fmt"

// 结构体构造函数：构造一个结构体实例的函数
// 结构体是值类型，每次进行实例化时都会进行值拷贝

type person struct {
	name, city string
	age		   int8
}

// 值拷贝的话返回的是一个结构体，可能开销很大
// 使用构造函数来返回一个结构体的指针，则可以节省开销（因为内存地址是固定长度）
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}


func main() {
	p1 := newPerson("小王", "北京", int8(18))
	fmt.Printf("type:%T value:%#v\n", p1, p1)
}