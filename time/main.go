package main

// time package demo

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now() // 得到一个时间对象
	fmt.Println(now)

	after1hour := now.Add(time.Hour)
	fmt.Println(after1hour)

	duration := after1hour.Sub(now)
	fmt.Println(duration)
	
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
	fmt.Println("over")

	// 定时器
	// for tmp := range time.Tick(2 * time.Second) {
	// 	fmt.Println(tmp)
	// }

	// 时间格式化 Y     m  d   H  M  S
	//           2006  01 02  15 04 05
	ret1 := now.Format("2006.01.02 15:04:05")	
	fmt.Println(ret1)
	
	// 解析 字符串类型的时间
	timeStr := "2019/08/07 15:00:00"
	// 1.拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return 
	}

	// 2. 根据时区去解析一个字符串类型的时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil{
		fmt.Printf("parse timeStr failed. err: %v\n", err)
		return 
	}
	fmt.Println(timeObj)


}