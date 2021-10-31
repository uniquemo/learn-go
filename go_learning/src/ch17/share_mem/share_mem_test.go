package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}

	time.Sleep(1 * time.Second) // 为了让主线程等待所有的goroutine执行完毕，否则无法获取到正确的结果
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1) // 增加一个goroutine
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 完成一个goroutine
		}()
	}

	wg.Wait() // 等待所有的goroutine执行完毕
	t.Logf("counter = %d", counter)
}
