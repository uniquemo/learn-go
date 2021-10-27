package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	var i int = 0
	for i < 5 {
		t.Log(i)
		i++
	}
}
