package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // 与数组的区别：没有声明长度，因为是可变长的
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)    // 初始化长度为3，容量为5
	t.Log(len(s2), cap(s2))    // 3 5
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) // [Jun Jul Aug] 3 7
	summer[0] = "Unknow"
	t.Log(Q2)   // [Apr Unknow Jun]
	t.Log(year) // [Jan Feb Mar Apr Unknow Jun Jul Aug Sep Oct Nov Dec]
}

// 切片不可以进行比较
func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	t.Log(a == b)
	b[0] = 2
	t.Log(a == b)
}
