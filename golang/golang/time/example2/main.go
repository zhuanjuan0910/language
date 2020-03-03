package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间戳
	t := time.Now().Unix()
	fmt.Println(t)
	//格式化当前时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//时间戳转string格式化时间
	time1 := time.Unix(0, 0).Format("2006-01-02  15:04:05")
	time2 := time.Unix(1582856857, 0).Format("2006年01月02日  15时04分05秒")
	fmt.Println(time1, time2)
	// str格式化时间转时间戳 
    // 方法一   2018-03-30 15:24:59
    the_time := time.Date(2018, 3, 30, 15, 24, 59, 0, time.Local)
    unix_time := the_time.Unix()
    fmt.Println(unix_time)
    fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))
    // 方法二 , 使用time.Parse
    /*
        返回的不是本地时间, 而是 UTC , 会自动加8小时.
    */
    the_time, err := time.Parse("2006-01-02 15:04:05", "2018-03-30 15:24:59")
    if err == nil {
        unix_time := the_time.Unix()
        fmt.Println(unix_time)
        fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))		
    }
    // 使用time.ParseInLocation
    the_time, err = time.ParseInLocation("2006-01-02 15:04:05", "2018-03-30 15:24:59",time.Local)
    if err == nil {
        unix_time := the_time.Unix()
        fmt.Println(unix_time)
        fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))		
    }