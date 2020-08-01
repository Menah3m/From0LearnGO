package main

import (
	"fmt"
)

func main(){
	var ch1 chan int  //声明一个传递int类型数据的通道
	// var ch2 chan int
	ch1 = make(chan int, 1)// 带缓冲区通道 -> 异步通道
	// ch2 = make(chan int)  //  无缓冲区通道 -> 同步通道
	ch1 <- 10
	// ch2 <- 100       // 因为没有缓冲区，所以无法将值放进去，就会等待，从而造成deadlock
	x := <-ch1
	// y := <-ch2
	fmt.Println(x)
	// fmt.Println(y)
	// len(ch1)  //取channel中元素个数
	// cap(ch1)  //取channel的容量
	close(ch1)
	// close(ch2)

}