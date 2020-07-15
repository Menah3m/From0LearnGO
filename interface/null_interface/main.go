package main

import (
	"fmt"
)


// 空接口
// 接口中没有定义任何方法时，该接口就是一个空接口
// 任意类型都实现了空接口 --> 空接口变量可以存储任意值
type xxx interface{    

}

// 空接口的应用
// 1. 空接口类型作为函数的参数
// 2. 空接口可以作为map的value

func main(){
	var x interface{} // 定义一个空接口变量
	x = "hello"       // x在内存中保存方式： 动态类型+动态值
	// fmt.Println(x)
	x = 100
	// fmt.Println(x)
	x = false
	// fmt.Println(x)

	// var m = make(map[string]interface{}, 16) // 空接口作为map的value
	// m["name"] = "张三"
	// m["age"] = 18
	// m["hobby"] = []string{"篮球","足球","双色球"}
	// fmt.Println(m)

	ret, ok := x.(string)   //类型断言:猜的不对时 ok=false, ret=string类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Printf("是字符串类型:",ret)
	}
	
	// 使用switch进行类型断言
	switch t := x.(type){
	case string:
		fmt.Printf("是字符串类型,value:%v\n",t)
	case bool:
		fmt.Printf("是布尔类型,value:%v\n", t)
	case int:
		fmt.Printf("是int类型,value:%v\n", t)
	default:
		fmt.Println("猜不到,value:%v\n", t)
	}


}