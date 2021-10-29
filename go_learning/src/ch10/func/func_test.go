package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	timeSpent(slowFun)(10)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	sum := Sum(1, 2, 3, 4)
	t.Log(sum)
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDerfer(t *testing.T) {
	defer Clear()
	defer func() {
		t.Log("defer 1")
	}()
	defer func() {
		t.Log("defer 2")
	}()
	fmt.Println("hello")
	panic("err") // 出错后，defer代码仍能执行，非defer代码则终止执行了
	fmt.Println("End")
}
