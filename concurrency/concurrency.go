package concurrency

import (
	"fmt"
	"time"
	"runtime"
)

func Go()  {
	fmt.Println("Go Go Go!!!")
}

func testGo ()  {
	fmt.Println("1")
	go Go() //异步执行
	fmt.Println("2")
	fmt.Println("ok")
	time.Sleep(5*time.Second) //暂停2秒
	fmt.Println("123")
	/**
	1
	2
	ok
	Go Go Go!!!
	123
	 */
}

// 函数Publish在给定时间过期后打印text字符串到标准输出
// 该函数并不会阻塞而是立即返回
func publish(text string, delay time.Duration)  {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

func testpublish()  {
	publish("A goroutine starts a new thread of execution.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// 等待发布新闻
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")
}

func goprocs()  {
	num := runtime.NumCPU()    // 本地机器的逻辑CPU个数
	runtime.GOMAXPROCS(num)    // 设置可同时执行的最大CPU数，并返回先前的设置
	fmt.Println(num)
}

func init()  {
	//testGo()
	goprocs()
}