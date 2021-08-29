package main

import "fmt"

func main() {
	fmt.Println("sum :", 3+2)   // sum - int
	fmt.Println("sum :", 2+3.5) // sum - float64
	fmt.Println("dif :", 3-1)   // difference - int
	fmt.Println("dif :", 3-0.5) // difference - float64
	fmt.Println("prod:", 4*5)   // product - int
	fmt.Println("prod:", 5*2.5) // product - float64
	fmt.Println("quot:", 8/2)   // quotient - int
	fmt.Println("quot:", 8/1.5) // quotient - float64

	// remainder is only for integers
	fmt.Println("rem :", 8%3)
	// fmt.Println("rem:", 8.0%3) // error

	// you can do this
	// since the fractional part of a float is zero
	f := 8.0
	fmt.Println("rem :", int(f)%3)

	floatExample()
}

func floatExample() {
	// what's the value of the ratio?
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio) // 1

	// explain
	// above expression equals to this:
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio) // 1

	// how to fix it?
	//
	// remember, when one of the values is a float value
	// the result becomes a float
	ratio = float64(3) / 2
	fmt.Println(ratio) // 1.5

	// or
	ratio = 3.0 / 2
	fmt.Println(ratio) // 1.5
}
