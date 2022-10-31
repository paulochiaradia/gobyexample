package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"          // literal value representing the word “hello” in the Thai language
	fmt.Println("len:", len(s)) //length of the raw bytes stored within

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println() //This loop generates the hex values of all the bytes that constitute the code points in s.

	fmt.Println("Rune count :", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	} //A range loop handles strings specially and decodes each rune along with its offset in the string.

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// This demonstrates passing a `rune` value to a function.
		examineRune(runeValue)

	}
}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
