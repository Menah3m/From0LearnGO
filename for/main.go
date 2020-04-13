package main

import "fmt"

// for循环

func main() {

	// 1. 最基本样式
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. 可以省略 初始语句 ，但是必须保留初始语句后面的";"
	j := 5
	for ; j > 0; j-- {
		fmt.Println(j)
	}

	// 3. 省略初始语句和结束语句,相当于其他语言的while语句
	y := 5
	for y > 0{
		fmt.Println(y)
		y--
	}

	// 4. 无限循环
	// for {
	// 	fmt.Println("Hello,world~")
	// }

	// 5. break   作用：跳出for循环

	for b:=0; b < 5; b++ {
		fmt.Println(b)
		if b == 3 {
			break
		}
	}

	// 6. continue 作用：跳过本次的循环，继续下一次循环

	for c :=0; c < 5; c++ {
		if c == 3 {
			continue
		}
		fmt.Println(c)
		
	}

	// 7. for range 循环  作用：用于遍历数组、切片、字符串、map及channel
}