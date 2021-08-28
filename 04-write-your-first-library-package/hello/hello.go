package main

import (
	"fmt"
	"log"

	"example.com/printer"
)

func main() {
	callHello()
	callHellos()
}

func callHello() {
	log.SetPrefix("printer: ")
	log.SetFlags(0)

	message, err := printer.Hello("momo")

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func callHellos() {
	names := []string{"Mary", "Henry", "Cherry"}

	message, err := printer.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
