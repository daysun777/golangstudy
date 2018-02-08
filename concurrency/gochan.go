package concurrency

import "fmt"

func testchan() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	fmt.Println(<-ch) /// 输出字符串"Hello!"
	fmt.Println(<-ch)	// 输出零值 - 空字符串""，不会阻塞
	fmt.Println(<-ch) // 再次打印输出空字符串""
	v, ok := <-ch
	fmt.Println(v)	// 变量v的值为空字符串""
	fmt.Println(ok) //变量ok的值为false
}

func init1() {
	testchan()
}
