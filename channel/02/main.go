package main

import (
	"fmt"
)

/*
两个goroutine
	1.生成0-100的数字发送到ch1
	2.从ch1中取出数字并计算它的平方，把结果给ch2

*/
func f1(ch chan<- int)  { //定义ch为 只写channel
	for i:=0;i<100;i++{
		ch <- i
	}
	
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	// 从channel中取值的方式1： ok法
	for {
		tmp, ok := <- ch1
		if !ok{
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main()  {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
    go f2(ch1, ch2)
	
	// 从channel中取值的方法2： for range
	for ret := range ch2{
		fmt.Println(ret)	
	}
}