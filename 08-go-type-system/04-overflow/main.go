package main

import (
	"fmt"
	"math"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	var (
		width  uint8 = 255
		height       = 255 // int
	)

	width++

	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height)
}

func test2() {
	// go catches overflow at compile-time
	//
	// fmt.Println(int8(math.MaxInt8 + 1)) // overflows

	// but it cannot catch them in runtime
	var n int8 = math.MaxInt8

	// wrap arounds to its negative maximum
	fmt.Println("max int8     :", n)   // 127
	fmt.Println("max int8 + 1 :", n+1) // -128

	// wrap arounds to its positive maximum
	n = math.MinInt8
	fmt.Println("min int8     :", n)   // -128
	fmt.Println("min int8 - 1 :", n-1) // 127

	// wrap arounds to its maximum
	var un uint8

	fmt.Println("max uint8    :", un)   // 0
	fmt.Println("max uint8 - 1:", un-1) // 255

	// wrap around to its minimum
	un = math.MaxUint8
	fmt.Println("max uint8 + 1:", un+1) // 0

	// floats goes to infinity when overflowed
	f := float32(math.MaxFloat32)
	fmt.Println("max float    :", f*1.2)
}

func test3() {
	// uint16 max value is 65535
	big := uint16(65535)

	// uint8 destroys its value
	// to its own max value which is 255
	//
	// 65535 - 255 is lost.
	small := uint8(big)

	// fmt.Printf("%b %d\n", big, big)
	// fmt.Printf("%b %[1]d\n", big)

	fmt.Printf("%016b %[1]d\n", big)
	fmt.Printf("%016b %[1]d\n", small)
}
