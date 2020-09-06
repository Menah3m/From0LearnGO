package main

// time package demo

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now() // 得到一个时间对象
	fmt.Println(now)
	
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	// 时间戳：从1970.1.1 到现在的一个总秒数
	timestamp := now.Unix()
	timestamp1 := now.UnixNano()
	fmt.Println(timestamp, timestamp1)

	// 将时间戳 转换为 具体的时间格式
	// 1599371075
	t := time.Unix(1599371075, 0)
	fmt.Println(t)


	// 时间间隔 duration
	// time.Duration类型
	
	n := 5  //type int
	time.Sleep(time.Duration(n)*time.Second)
	println("over")
}