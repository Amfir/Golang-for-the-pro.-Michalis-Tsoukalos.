package main

import (
	"fmt"
	"unicode"
)

func main() {
	const s = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"

	for i := 0; i < len(s); i++ {
		if unicode.IsPrint(rune(s[i])){
			fmt.Printf("%c\n", s[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}