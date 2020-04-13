package main

import "fmt"

// Go语言中的运算符

func main() {
	// 1. 算数运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(5 / 2)  //2 
	fmt.Println(5 % 2)  //1 求余
	a++                 // a++ 是一个语句，不是运算符，后面不接数字
	fmt.Println(a)

	// 2. 关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 == 5)

	// 3. 逻辑运算符  &&  ||  !
	fmt.Println(10>5 && 2==1)
	fmt.Println(10>5 || 2==1)
	fmt.Println(!(10>5))

	// 4. 位运算符  
	// & 二进制位 相与   
	// | 二进制位 相或
	// ^ 二进制位 异或  （不一样则为1）
	// << 左移n位  也就是乘以2的n次方
	// >> 右移n位  也就是除以2的n次方

	b1 := 1 // 001
	b2 := 5 // 101

	fmt.Println(b1 & b2) // 001 -> 1
	fmt.Println(b1 | b2) // 101 -> 5
	fmt.Println(b1 ^ b2) // 100 -> 4
	fmt.Println(b1 << 3) // 1000 -》 1*2^3=8 (高位丢弃，低位补0)
	fmt.Println(b2 >> 1) // 010 -> 5/2^1 = 2（低位丢弃，高位补0)
}