package main

import (
	"fmt"
	"sync"
	"time"
)


// 读写互斥锁

/*
互斥锁是完全互斥的，很多实际场景下读操作比写操作多，当并发地读取资源时并不需要加锁，因此完全互斥会造成资源的浪费
读写互斥锁是更好的选择。

读写锁分为：读锁和写锁。
当一个goroutine获得读锁时，如果另一个goroutine申请读锁，则会直接获得锁，如果申请写锁，则会等待读操作完成。
当一个goroutine获得写锁时，则不管另一个goroutine申请读锁还是写锁，都要进行等待

下面的demo是使用互斥锁和读写互斥锁下执行读写操作
所需时间对比：
使用【互斥锁】 1.2846508s
使用【读写锁】 108.306ms

*/

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func read()  {
	// lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
	}
	


func write()  {
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond*10)
	// lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}

func main()  {
	start := time.Now()
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}

	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
