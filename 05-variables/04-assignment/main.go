package main

import "fmt"

func main() {
	var speed int
	fmt.Println(speed)

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	var force float64
	// Go automatically converts the typeless integer literal to float64 automatically
	force = 1
	_ = force

	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	// mutiple assignment
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	// mutiple short func
	_, b := multi()
	fmt.Println(b)

	// swapper
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)
}

func multi() (int, int) {
	return 5, 4
}
