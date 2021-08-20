package main

import "fmt"

func nope() {
	const ok = true
	var hello = "Hello"

	_ = hello
}

func main() {
	// ERROR:
	fmt.Println(hello, ok)
}
