package main

import "fmt" // file scope

// package scope
const ok = true

// package scope
func main() { // block scope starts
	var hello = "Hello"
	fmt.Println(hello, ok)
} // block scope ends
