package main


import (
	"fmt"
	"sync"
)

// sync.Map

/*
Go语言中原生的map不支持并发写操作
可以用sync.Map实现并发写入
sync.Map不需要进行初始化就可以使用
*/

var (
	m = make(map[int]int)
	wg sync.WaitGroup
	m2 = sync.Map{}
)

func get(key int)int  {
	return m[key]
}

func set(key int, value int)  {
	m[key]=value
}

// 原生map进行并发操作会报错：【fatal error: concurrent map writes】
// func main()  { 
// 	for i:=0;i<20;i++{
// 		wg.Add(1)
// 		go func(i int){
// 			set(i, i+100)
// 			fmt.Printf("key:%v, value:%v", i, get(i))
// 			wg.Done()
// 		}(i)
		
// 	}
// 	wg.Wait()
	
// }

//使用sync.Map进行并发操作
func main()  {
	for i:=0;i<20;i++{
		wg.Add(1)
		go func(i int){
			m2.Store(i, i+100)  //设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("key:%v, value:%v\n", i, value)
			wg.Done()
		}(i)
		
	}
	wg.Wait()
	
}