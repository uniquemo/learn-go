package panic_recover_test

import (
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 恢复错误
			t.Log("panic recover:", err)
		}
	}()
	fmt.Println("Start")
	panic("panic")
	// os.Exit(-1)
	// fmt.Println("End")
}
