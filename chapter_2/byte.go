package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("A String")
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Printf("%s\n", b)
	fmt.Println()

	// руна
	r := '$'
	fmt.Println(r)
	fmt.Printf("%c\n\n", r)

	aString := "Hello World! $"
	fmt.Println("first character", string(aString[0]))
	fmt.Println()

	// руна
	r = '$'
	fmt.Println("As an int32 value:", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	fmt.Println()

	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	s := 100
	fmt.Println(string(s))
	fmt.Println(strconv.Itoa(s))
	fmt.Println(strconv.FormatInt(int64(s), 10))
}
