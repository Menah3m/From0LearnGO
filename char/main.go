package main

import "fmt"


// 字符 组成字符串的元素称为字符
func main() {
	// byte  uint8的别名 表示ASCII码
	// rune  int32的别名 表示中文，日文，韩文等时 使用rune
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

}