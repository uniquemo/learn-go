package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"

	// strings are made up of bytes
	// len function counts the bytes in a string value
	fmt.Println(len(name))

	unicodeLen()
}

func unicodeLen() {
	// strings are made up of bytes

	// len function counts the bytes in a string value.
	//
	// This string literal contains unicode characters.
	//
	// And, unicode characters can be 1-4 bytes.
	// So, "İnanç" is 7 bytes long, not 5.
	//
	// İ = 2 bytes
	// n = 1 byte
	// a = 1 byte
	// n = 1 byte
	// ç = 2 bytes
	// TOTAL = 7 bytes
	name := "İnanç"
	fmt.Printf("%q is %d bytes\n", name, len(name))

	// To get the actual characters (or runes) inside
	// a utf-8 encoded string value, you should do this:
	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))
}
