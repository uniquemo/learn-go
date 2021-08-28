package main

import "fmt"

var isLiquid bool // package variable

func main() {
	var speed int
	fmt.Println(speed) // 0

	// CORRECT DECLARATIONS
	var SpeeD int
	// underscore is allowed but not recommended
	var _speed int
	// Unicode letters are allowed
	var 速度 int
	// let's assign the variable to the blank-identifier
	// so that Go compiler won't get grumpy
	_, _, _, _ = speed, SpeeD, _speed, 速度

	// int
	var nFiles int
	var counter int
	var nCPU int
	fmt.Println( // 0 0 0
		nFiles,
		counter,
		nCPU,
	)

	// float64
	var heat float64
	var ratio float64
	var degree float64
	fmt.Println( // 0 0 0
		heat,
		ratio,
		degree,
	)

	// bool
	var off bool
	var valid bool
	var closed bool
	fmt.Println( // false false false
		off,
		valid,
		closed,
	)

	// string
	var msg string
	var name string
	var text string
	fmt.Println(
		msg,
		name,
		text,
	)
	fmt.Printf("%q\n", text)

	// multiple declarations
	var (
		speed2 int
		heat2  float64
		off2   bool
		brand2 string
	)
	fmt.Println(speed2)
	fmt.Println(heat2)
	fmt.Println(off2)
	fmt.Printf("%q\n", brand2)

	// parallel declarations
	var speed3, velocity3 int
	fmt.Println(speed3, velocity3)

	var firstName, lastName string = "", ""
	fmt.Printf("%q %q\n", firstName, lastName)
}
