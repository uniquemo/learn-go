package main

func main() {
	// aliased types are the same types
	// just with different names
	// for readability and refactoring
	var (
		// type byte = int8
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	uint8Val = byteVal // ok

	var (
		// type rune = int32
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val // ok

	runeVal = rune(int32Val)
	runeVal = rune(runeVal)

	// keep the compiler happy
	_, _, _, _ = byteVal, uint8Val, intVal, runeVal
}
