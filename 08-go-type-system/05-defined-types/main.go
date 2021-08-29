package main

import (
	"fmt"
	"time"
)

// type definitions usually declared at the package level
//
// EXERCISE: Move the declaration into main()'s scope
//
type (
	gram  float64 // float64 is the underlying type of gram
	ounce float64 // float64 is the underlying type of ounce
	Gram  gram
)

// above code is the same as the following:
// type gram  float64
// type ounce float64

func main() {
	test1()
	test2()
}

func test1() {
	h, _ := time.ParseDuration("4h30m")

	// why would you want to create a new type?

	// 1- adding new methods to the type
	fmt.Println(h.Hours(), "hours")

	// 2- make it a distinct type for type-safety
	// you can't use the defined type
	//   with its underlying type directly.
	//
	// you need to convert one of them.
	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	// type of `h` and its underlying type are different
	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))
}

func test2() {
	// type definitions are also allowed in blocks

	// underlying types are the same
	var g gram = 1000 // gram  -> float64
	var o ounce       // ounce -> float64
	var salt Gram = 100

	// when the underlying types are the same
	// then they can be converted between:
	// gram, ounce or float64

	// afterwards, you'll see also that,
	// the important thing is the structure
	// of the type. not just its name.
	//
	// float64 has the real structure, representation,
	// and size.
	//
	// so, it gives the structure to the newly defined type.

	// TYPE ERROR: ounce and grams are different types
	// o = g * 0.035274

	// BUT: They're convertable to each other
	o = ounce(g) * 0.035274

	fmt.Printf("%g grams is %.2f ounce\n", g, o)
	fmt.Printf("Type of salt: %T\n", salt)
}
