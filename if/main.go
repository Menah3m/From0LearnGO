package main

import "fmt"

// if判断
func main() {
	// 1. 基本写法
	// var score int = 85
	// if score > 90 {
	// 	fmt.Println("A")
	// } else if score > 80 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	// 2. 特殊写法：如果有变量只在if结构里被用到，
	//             则可以在if语句开头声明它
	
	if score := 85; score > 90 {
		fmt.Println("A")
	} else if score > 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 3. 省略写法： if语句中，if、else if、else不需要全部都有
	//              if是必须的，其余的可以省略
	score := 100
	if score > 95 {
		fmt.Println("A+")
	} else {
		fmt.Println("A")
	}

}