package channels

import (
	"fmt"
)

/**
关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以 用来给这个通道的接收方传达工作已经完成的信息。
 */

func closechan() {
	jobs := make(chan int, 5) //申请通道
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job: ", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done //阻塞
}

func rangchan() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"

	close(queue)
	fmt.Println(queue) //0xc04208a000
	for v := range queue {
		fmt.Println(v)
	}
}

func init() {
	closechan()
	rangchan()
}
