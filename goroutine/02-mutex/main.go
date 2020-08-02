package main

import (
	"fmt"
	"sync"
)

/*
当多个goroutine对同一个资源进行访问时，会出现竞态问题
可以使用互斥锁（Mutex）来解决
互斥锁可以保证同一时刻有且只有一个goroutine访问资源，其他goroutine进入临界区，进行等待
当互斥锁释放时，从等待的goroutine中随机选择一个goroutine分配互斥锁

*/

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add()  {
	for i:=0;i<5000;i++{
		lock.Lock()   // 加锁
		x = x + 1     // 操作
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main()  {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}