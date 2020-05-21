package main

import "fmt"

// 嵌套结构体的字段冲突: 如果产生了字段冲突，则需要指明对应的结构体

type Address struct{
	Provice    string
	City	   string
	UpdateTime string

}

type Email struct{
	Addr       string
	UpdateTime string
}

type Person struct{
	Name      string
	Age       int8
	Gender    string
	Address   // 嵌套另外一个结构体:两种方式：1 字段名+ 字段类型 2  仅有字段类型（此时为匿名字段）
	Email     // 嵌套另一个结构体

}


func main(){
	p := Person{
		Name:"李四",
		Age:18,
		Gender:"男",
		Address:Address{
			Provice:"山东",
			City:"威海",
			UpdateTime:"20200511",
		},
		Email:Email{
			Addr:"aaaaa@qq.com",
			UpdateTime:"20200520",
		},
	}

	fmt.Println("%#v\n", p)
	fmt.Println(p.Address.UpdateTime)
	fmt.Println(p.Email.UpdateTime)

}