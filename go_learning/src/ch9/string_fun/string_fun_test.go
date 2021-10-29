package string_fun_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa((10))
	t.Log("str" + s) // str10
	if n, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + n) // 20
	}
}
