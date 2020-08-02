package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //使用sync包创建wg实例

// func hello(i int){
// 	fmt.Println("hello,", i)
// 	wg.Done()   //计数器 -1
// }



// func main() {  // 开启一个主goroutine去执行main函数
	
// 	// for i:=0; i<10; i++{
// 	// 	wg.Add(1)  //计数器 +1
// 	// 	go hello(i) //开启了一个独立的goroutine去执行hello函数
// 	// }
// 	for i:=0; i<10; i++{
// 		wg.Add(1)
// 		go func(i int) { // 开启一个goroutine去执行匿名函数
// 			fmt.Println("hello",i)  
// 			wg.Done()
// 		}(i)
// 	}

// 	fmt.Println("hello, main!")

// 	wg.Wait()   // 阻塞，等待计数器归0 然后结束
// }

func a(){
	for i:=0; i<10; i++{
		fmt.Println("A:",i)
		
	}
	wg.Done()
}

func b(){
	for i:=0;i<10;i++{
		fmt.Println("B:",i)
		
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2) //占用cpu核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	
}