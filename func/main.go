package main

import "fmt"

// 函数
func sayHello(){
	fmt.Println("Hello,world~")
}

// 定义一个带参数的函数
func sayHelloTo(name string){
	fmt.Println("hello", name)
}

// 定义接收多个参数的函数 
//如果给返回值声明了一个变量，则函数体中直接使用这个变量即可
//如果返回值没有声明变量，只声明了类型，那么函数体中需要声明才能使用
func intSum(a int, b int) (int) {   
	ret := a + b
	return ret
}

// 函数接收可变参数   
// 函数如果接收多种参数，则可变函数一定要放在固定参数后面
func intSum3(a int, b... int) int {
	ret := a
	for _, v := range b {
		ret = ret + v
	}
	return ret
}


func main(){
	//函数调用
	sayHello()
	sayHelloTo("Tom")

	ret := intSum(2, 3)
	fmt.Println(ret)

	r1 := intSum3(0)        // a=0 b=[]
	r2 := intSum3(1)        // a=1 b=[]
	r3 := intSum3(1, 2)     // a=1 b=[2]
	r4 := intSum3(1, 2, 3)  // a=1 b=[2 3]

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

}

