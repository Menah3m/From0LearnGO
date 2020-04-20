package main

//函数进阶

import "fmt"

// 函数进阶-变量作用域

// 定义全局变量
var num = 10

func testGolbal() {
	// 可以在函数中直接访问全局变量，但是如果函数内有同名变量，则优先访问函数中定义的该变量
	// 可以理解成：先找自己有没有，没有找到则去外面找全局
	num := 100
	name := "张三"
	fmt.Println("全局变量：", num)
	fmt.Println(name)
}

func add(x, y int) (int) {
	return x + y
}

func calc(x, y int, op func(int, int)(int))(int) {
	return op(x , y)
}

func main() {
	// 函数可以作为变量
	abc := testGolbal   
	fmt.Printf("%T\n", abc)
	abc()

	// 函数可以作为参数
	r1 := calc(10, 20, add)
	fmt.Println(r1)



	// fmt.Println(name) // 外层无法直接访问函数内部的变量(局部变量)
	// for i := 0; i < 5; i++ { 
	// 	fmt.Println(i)
	// }
	// // fmt.Println(i)  // 控制语句块中定义的变量也无法被外部访问
   

	
}

