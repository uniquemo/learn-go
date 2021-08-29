package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	var n int

	// ALTERNATIVES:
	// n = n + 1
	// n += 1

	// BETTER:
	n++

	fmt.Println(n)
}

func test2() {
	n := 10

	// ALTERNATIVES:
	// n = n - 1
	// n -= 1

	// BETTER:
	n--

	fmt.Println(n)
}

// incdec is a statement
func test3() {
	var counter int

	// following "statements" are correct:
	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	fmt.Printf("There are %d line(s) in the file\n", counter)

	// the following "expressions" are incorrect:
	// counter = 5+counter--
	// counter = ++counter + counter--
}
