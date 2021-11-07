package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	var (
		uniqueCount int
		count int
	)
	wordExist := map[string]bool{}

	for in.Scan() {
		word := strings.ToLower(in.Text())
		if wordExist[word] == false {
			wordExist[word] = true
			uniqueCount++
		}
		count++
	}
	fmt.Printf("there are %d words, %d of them are unique", count, uniqueCount)
}