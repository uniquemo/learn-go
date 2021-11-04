package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // 发送一个空结构体消息
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCancelled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh: // 如果cancelCh有数据（收到一个消息），则返回true
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}

	// cancel_1(cancelChan) // 用这种方式会导致只有一个goroutine被取消，因为生产的消息只有一个
	cancel_2(cancelChan) // 用这种方式会导致所有的goroutine被取消

	time.Sleep(time.Second * 1)
}
