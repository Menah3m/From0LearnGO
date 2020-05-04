package main


import "fmt"

// 自定义类型和类型别名示例

// 1. 自定义类型

//MyInt 是基于int的自定义类型
type MyInt int

// 2. 类型别名（Go1.9版本新增）

// type byte = uint8
// type rune = int32

// NewInt 是int的类型别名
type NewInt = int 

func main(){
	var i MyInt
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)

	var x NewInt
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)

}