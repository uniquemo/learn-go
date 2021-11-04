package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "[Service] Done"
}

func AsyncService() chan string {
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
	// 多路选择
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 1000):
		t.Error("time out")
	}
}
