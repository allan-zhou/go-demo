package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Now()
	fmt.Println(t) // 2018-01-11 15:12:29.3918782 +0800 CST m=+0.00198140
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d.%d  \n",t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())  //2018-01-11 15:31:31.875685900

	t = time.Now().UTC()
	fmt.Println(t)  //2018-01-11 07:15:35.5651095 +0000 UTC
	t = time.Now().Local()
	fmt.Println(t) //2018-01-11 15:16:29.1261032 +0800 CST

	fmt.Println(t.Format(time.ANSIC))     //Thu Jan 11 15:21:35 2018
	fmt.Println(t.Format(time.UnixDate))  //Thu Jan 11 15:21:35 CST 2018

	// Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。
	// 纳秒：一秒的十亿分之一，即等于10的负9次方秒（1 ns = 10-9 s）。
	duration :=  time.Duration(2) //2ns
	fmt.Printf("duration: %d  %s \n", duration, duration)

	start := time.Now()
	duration =  time.Duration(2) * time.Second //2s
	fmt.Printf("duration: %d  %s \n", duration, duration)
	time.Sleep(duration) 
	end := time.Now()
	elasped := end.Sub(start)
	fmt.Printf("elasped:%s \n",elasped)
	
}