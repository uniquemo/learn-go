package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrLessThanTwo = errors.New("n should be not less than 2")
var ErrLargerThanHundred = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrLessThanTwo
	}

	if n > 100 {
		return nil, ErrLargerThanHundred
	}

	fibList := []int{0, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	var fibList []int
	var err error

	if fibList, err = GetFibonacci(-10); err != nil {
		if err == ErrLessThanTwo {
			fmt.Println("It is less.")
		}

		if err == ErrLargerThanHundred {
			fmt.Println("It is larger.")
		}

		t.Error(err)
		return
	}

	t.Log(fibList)
}
