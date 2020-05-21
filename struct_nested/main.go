package main

import "fmt"

// 结构体的匿名字段：省略结构体中字段的名称，只保留类型(类型要在结构体中唯一)


type Person struct{
	string
	int8

}

// 嵌套结构体:结构体字段太多的时候可以把一部分单独出来作为一个结构体

// type Person2 struct{
// 	name     string
// 	age      int8
// 	Gender   string
// 	provice  string
// 	city     string
// }
type Address struct{
	Provice    string
	City	   string
}

type Person2 struct{
	Name      string
	Age       int8
	Gender    string
	Address   // 嵌套另外一个结构体:两种方式：1 字段名+ 字段类型 2  仅有字段类型（此时为匿名字段）
}


func main(){
	p1  := Person{
		string:"张三",
		int8:18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int8)

	p2 := Person2{
		Name:"李四",
		Age:19,
		Gender:"男",
		Address:Address{
			Provice:"北京",
			City:"北京",
		},

	}

	fmt.Println(p2)
	fmt.Println(p2.Address)
	fmt.Println(p2.Address.Provice)
	fmt.Println(p2.Provice)   //匿名字段可以直接被访问
	
}