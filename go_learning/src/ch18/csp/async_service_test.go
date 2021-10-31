package async_service_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "[Service] Done"
}

func otherTask() {
	fmt.Println("[Other] working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("[Other] Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// retCh := make(chan string) // 创建一个channel，会阻塞
	retCh := make(chan string, 1) // 创建一个channel，并设置缓冲大小为1

	go func() { // 开启一个goroutine
		ret := service()
		fmt.Println("[Groutime] returned result.")
		retCh <- ret // 将结果发送到channel中
		fmt.Println("[Groutine] service exited.")
	}()

	return retCh // 返回channel
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

/*

buffer channel的执行结果：

[Other] working on something else
[Groutime] returned result.
[Groutine] service exited.
[Other] Task is done.
[Service] Done

非buffer channel的执行结果：

[Other] working on something else
[Groutime] returned result.
[Other] Task is done.
[Service] Done
[Groutine] service exited.

*/
