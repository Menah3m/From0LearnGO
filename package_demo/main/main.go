package main

import (
	jiafa "github.com/Menah3m/From0LearnGo/package_demo/calc" //当代码都放在$GOPATH目录下的话
	// 在包的路径前面可以设置包的别名 防止包名冲突
	// 我们导包的路径要从$GOPATH/src/后面继续写起
	"fmt"
)

func main(){
	fmt.Println("hello~")
	x := 1
	y := 2
	res := jiafa.Add(x, y)  //使用包的别名
	fmt.Println(res)
	fmt.Println(jiafa.Name)
}

func init(){
	fmt.Println("main.init()")
}