package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name string
	Sex  int
}

// Teacher结构体有了String()，所以该结构体指针的数组打印时会调用String()
func (t *Teacher) String() string {
	return fmt.Sprintf("{Name: %s, Sex: %d}", t.Name, t.Sex)
}

func main() {
	std1 := &Student{Name: "studen1", Age: 18}
	arr1 := []*Student{
		std1,
		{Name: "studen2", Age: 16},
	}
	fmt.Printf("studen1 => %v\n", std1)
	fmt.Printf("&studen1 => %v\n", &std1)
	fmt.Printf("结构体指针数组 arr1 => %v\n", arr1)

	arr2 := []*Teacher{
		{Name: "Teacher1", Sex: 0},
		{Name: "Teacher2", Sex: 1},
	}
	fmt.Printf("结构体指针数组 arr2 => %v\n", arr2)
}

/*

输出结果：

studen1 => &{studen1 18}
&studen1 => 0xc00018e018
结构体指针数组 arr1 => [0xc000198000 0xc000198018]
结构体指针数组 arr2 => [{Name: Teacher1, Sex: 0} {Name: Teacher2, Sex: 1}]

*/
