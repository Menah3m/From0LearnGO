package main

import "fmt"
//变量
func main(){
	//标准声明格式： var 变量名 变量类型
	var name string
	var age int
	var isOK bool
	fmt.Println(name, age, isOK)

	//批量声明格式
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	//声明变量同时指定初始值  格式：var 变量名 变量类型 = 变量值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	

	//类型推导   根据变量设置的初始值来判断变量类型
	var name2, age2 = "小小王子", 18
	fmt.Println(name2, age2)
	var name3 = "小王子"
	var age3 = 18
	fmt.Println(name3, age3)

	//简短变量声明：只能在函数内部使用
	m := 10
	fmt.Println(m)

	//匿名变量：用"_" 表示， 不占用内存空间，也不会分配内存（所以不存在重复声明）

	


}