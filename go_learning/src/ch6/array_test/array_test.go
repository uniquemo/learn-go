package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[1] = 5
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1, arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, val := range arr3 {
		t.Log(idx, val)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[1:3]
	arr3_sec[0] = 100
	t.Log(arr3_sec) // [100, 3]
	t.Log(arr3)     // [1, 100, 3, 4, 5]
}
