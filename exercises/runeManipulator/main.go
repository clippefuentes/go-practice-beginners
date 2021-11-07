package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}


	for _, word := range words {
	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("%s: Bytes is %d and rune length %d \n", word, len(word), utf8.RuneCountInString(word))

	// Print the bytes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings as rune literals
	// Hint: Use for range

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString
	_, first := utf8.DecodeRuneInString(word)
	_, second := utf8.DecodeLastRuneInString(word[first:])
	fmt.Printf("%s, first len: %d second len: %d \n", word[:first+second], first, second)
	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString

	// Slice and print the first two runes of the strings

	// Slice and print the last two runes of the strings

	// Convert the string to []rune
	// Print the first and last two runes
	}
}