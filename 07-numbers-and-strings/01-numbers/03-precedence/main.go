package main

import "fmt"

func main() {
	test1()
	test2()
}

func test1() {
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2), // same as above
	)

	fmt.Println(
		1+4-2,
		(1+4)-2, // same as above
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)
}

func test2() {
	n, m := 1, 5

	fmt.Println(2 + 1*m/n)
	fmt.Println(2 + ((1 * m) / n)) // same as above

	// let's change the precedence using parentheses
	fmt.Println(((2 + 1) * m) / n)
}
