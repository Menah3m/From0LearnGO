package main

import (
	"github.com/Menah3m/From0LearnGo/package_demo/calc" //当代码都放在$GOPATH目录下的话
						// 我们导包的路径要从$GOPATH/src/后面继续写起
	"fmt"
)

func main(){
	fmt.Println("hello~")
	x := 1
	y := 2
	res := calc.Add(x, y)
	fmt.Println(res)
}