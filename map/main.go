package main

import "fmt"

//map(映射)
func main() {
	// 光声明map类型，没有初始化，则是一个零值 nil
	var a map[string]int
	fmt.Println(a == nil)

	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// map中添加k-v对
	a["王二"] = 100
	a["刘三"] = 102
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type：%T\n", a)

	// 声明map的同时完成初始化
	b := map[int]bool{1:true, 2:false, 3:true}
	fmt.Printf("%#v\n", b)
	fmt.Printf("Type: %T\n", b)
}




































































