package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int){
	for job := range jobs{
		fmt.Printf("worker:%d, start job:%d\n", id, job)
		results <- job
		time.Sleep(time.Millisecond*500)
		fmt.Printf("worker:%d, start job:%d\n", id, job)
	}
	close(results)
}

func main()  {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for j:=0; j<3; j++{
		go worker(j, jobs, results)
	}

	// 发送5个任务
	for i:=0; i<5; i++{
		jobs <- i
	}
	close(jobs)
	// 输出结果
	for ret := range results{
		fmt.Println(ret)
	}

}