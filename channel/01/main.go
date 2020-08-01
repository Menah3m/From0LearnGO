package main

import (
	"fmt"
)

func main(){
	var ch1 chan int  //声明一个传递int类型数据的通道
	ch <- 10
	fmt.Println(ch)
}