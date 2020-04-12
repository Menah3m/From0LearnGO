package main

import (
	"fmt"
	"math"
)
//基本数据类型
//需要输出格式化数据的时候 用 printf
//输出变量或者整型的时候 用 println
func main() {
	//十进制打印为二进制
	var n int = 10
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)

	//八进制 0开头表示
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)

	//十六进制 0x开头
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)

	// uint8 无符号整数，范围0-255 超过的话会出错
	var age uint8 = 100
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 布尔值
	var flag bool
	fmt.Println(flag)
	flag = true
	fmt.Println(flag)

	// 字符串
	var s1 string = "hello Beijing"
	var s2 string = "你好，北京"
	fmt.Println(s1)
	fmt.Println(s2)

	//打印windows下的一个路径  c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")

	// 转义字符
	fmt.Println("\t 制表符 \n 换行符")

	// 多行字符串
	s3 := `
	多行字符串
	--在两个反引号
	中间的字符会保持原来的样式输出
	比如\t  \n
	`
	fmt.Println(s3)



}