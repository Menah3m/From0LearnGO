package main

import "fmt"

func main(){
	var a *int   //只声明了指针，并没有初始化
	a = new(int) //在内存中初始化a，赋0值，返回指针
	fmt.Println(*a)
	*a = 100     
	fmt.Println(*a)

	var b map[string]int //使用make进行初始化，返回对应的类型
	b = make(map[string]int, 10)
	b["张三"] = 100
	fmt.Println(b)
}