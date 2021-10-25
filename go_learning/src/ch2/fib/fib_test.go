package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Println("hello")
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Error("test error")
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
