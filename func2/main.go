package main

//函数进阶

import (
	"fmt"
	"strings"
)

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

// 闭包
// 定义一个函数 它的返回值是一个函数
func a(name string) func() {
	
	return func(){
		fmt.Println("a内部定义的匿名函数：hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}

func calculate(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
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
   
	// 匿名函数和闭包
	// 匿名函数的定义方法：没有函数名的普通函数就是匿名函数
	// 匿名函数有两种使用方法：
	// 1. 将匿名函数赋值给一个变量，则可以用该变量代替这个匿名函数执行
	sayniming := func() {
		fmt.Println("匿名函数")
	}
	sayniming()

	// 2. 在定义的匿名函数最后加括号即可直接执行
	func() {
		fmt.Println("直接执行匿名函数")
	}()


	// 闭包 = 函数 + 外层变量的引用
	r := a("张二")  //r此时相当于一个闭包
	r()       // 相当于执行了函数a内部的匿名函数

	// 闭包的使用例子1
	re := makeSuffixFunc(".exe")
	rere := re("ex")
	fmt.Println(rere)

	re1 := makeSuffixFunc(".avi")
	rere1 := re1("ex")
	fmt.Println(rere1)
	
	// 闭包的使用例子2
	addRe, subRe := calculate(5)
	r2add := addRe(4)    // base = 5 + 4 = 9
	fmt.Println(r2add)
	r2sub := subRe(4)    // base = 9 - 4 = 5
	fmt.Println(r2sub)

}

