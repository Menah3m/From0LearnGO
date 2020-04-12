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

	// 普通的for循环遍历中文等特殊字符时会产生乱码
	s := "hello 北京"
	for i:=0;i<len(s);i++{
		fmt.Printf("%c\n", s[i])
	}
	// 处理中英文字符串遍历时使用range循环
	for _,r := range s{
		fmt.Printf("%c\n", r)
	}
}