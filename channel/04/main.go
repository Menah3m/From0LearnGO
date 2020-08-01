package main

import (
	"fmt"
	
)



func main()  {
	ch := make(chan int, 1)
	for i:=0;i<10;i++{
		select{
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥都没干")
		}
	}
}
/*

i=0  此时channel中没有值，故不可读，选择分支case ch<-i并将i的值传到ch中
i=1  此时ch中有值，可读，选择分支case x:=<-ch 输出ch中读到的i值
i=2  ...
.
.
.
select语句类似于switch语句，选择符合条件的分支执行
如果有多个分支满足条件，则随机选择
*/