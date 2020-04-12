package main

import (
	"fmt"
	"strings"
)

func main() {
	// 求字符串的长度
	s1 := "hello"
	s2 := "你好beijing"
	s3 := "北京"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 拼接字符串 '+'
	fmt.Println(s1 + s2)
	s4 := fmt.Sprintf("%s-%s", s1, s3)
	fmt.Println(s4)

	// 字符串的分割 strings.Split(目标字符串，分割标志)
	s5 := "how do you do"
	fmt.Println(strings.Split(s5," "))
	fmt.Printf("%T\n", strings.Split(s5, " "))

	// 判断是否包含 strings.Contains(目标字符串，目标字符)
	fmt.Println(strings.Contains(s5, "do"))

	// 判断前缀 strings.HasPrefix(目标字符串, 目标前缀字符)
	fmt.Println(strings.HasPrefix(s5, "how"))

	// 判断后缀 strings.HasSuffix(目标字符串, 目标后缀字符)
	fmt.Println(strings.HasSuffix(s5, "how"))

	// 判断子串的位置 strings.Index(目标字符串，目标字符)
	fmt.Println(strings.Index(s5,"do"))

	// 判断子串最后出现位置 strings.LastIndex(目标字符串，目标字符)
	fmt.Println(strings.LastIndex(s5, "do"))

	// 	join
	s6 := []string{"how", "do", "you", "do"}
	fmt.Println(s6)
	fmt.Println(strings.Join(s6,"+"))


}