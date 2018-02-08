package util

import (
	"fmt"
	"time"
)

func init()  {
	p := fmt.Println

	now := time.Now()

	p(now)  //2018-02-07 10:40:03.7485567 +0800 CST m=+0.002000101

	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())

	//时间戳
	secs := now.Unix() //秒
	nanos := now.UnixNano() //纳秒
	millis := nanos /1000000  //没有直接获取毫秒，需要转换

	p(secs)		//1518057711
	p(nanos)	//1518057711354052000
	p(millis)	//1518057711354

	p(time.Unix(1518057711, 0)) //2018-02-08 10:41:51 +0800 CST
	p(time.Unix(0,nanos))

}