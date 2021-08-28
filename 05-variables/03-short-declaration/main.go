package main

import "fmt"

func main() {
	// = is the assignment operator
	// when used within a variable declaration, it
	// initializes the variable to the given value

	// OPTION #1 (option #2 is better)
	// var safe bool = true

	// OPTION #2
	// var safe = true

	// OPTION #3 - SHORT DECLARATION (BEST)
	safe := true
	fmt.Println(safe)

	// multiple short declarations
	a, c := 42, "good"
	fmt.Println(a, c)

	sum := 27 + 3.5
	fmt.Println(sum)

	// short discard
	on, _ := true, true
	fmt.Println(on)

	// redeclaration
	age, yourAge := 10, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)
}
