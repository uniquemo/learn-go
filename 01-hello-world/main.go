// package main is a special package
// it allows Go to create an executable file
package main

import "fmt"

// "func main" is special.
// Go has to know where to start
// func main creates a starting point for Go
// After compiling the code, Go runtime will first run this function
func main() {
	// View DOC: go doc -src fmt Println
	fmt.Println("Hi! I want to be a Gopher!", "abc")
	fmt.Print("bbb", "ccc")
	fmt.Println("bbb", "ccc")
}
