// https://docs.google.com/document/d/1PLJdD4POmdgmCvbs0VMAQ5UiSg4tDfq_lJ_Isa6hr3A/edit?usp=sharing

// Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter counts the number of words written to it.
type WordCounter int

// LineCounter counts the number of lines written to it.
type LineCounter int

// Counter interface represents a generic writer method for counting.
type Counter interface {
	Write(p []byte) (n int, err error)
}

// Write method for WordCounter that counts the number of words in p.
func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count) // Add count to WordCounter
	return count, scanner.Err()
}

// Write method for LineCounter that counts the number of lines in p.
func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count) // Add count to LineCounter
	return count, scanner.Err()
}

func main() {
	var wc WordCounter
	var lc LineCounter

	text := "Hello, world!\nThis is a test\nWith multiple lines."

	// Count words
	wordCount, err := wc.Write([]byte(text))
	if err != nil {
		fmt.Println("Error counting words:", err)
	}
	fmt.Printf("Word count: %d\n", wordCount)
	fmt.Printf("WordCounter value: %d\n", wc)

	// Count lines
	lineCount, err := lc.Write([]byte(text))
	if err != nil {
		fmt.Println("Error counting lines:", err)
	}
	fmt.Printf("Line count: %d\n", lineCount)
	fmt.Printf("LineCounter value: %d\n", lc)
}
