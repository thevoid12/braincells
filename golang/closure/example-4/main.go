package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// makeFilter returns a closure that checks if a line contains the keyword
func makeFilter(keyword string) func(string) bool {
	return func(line string) bool {
		return strings.Contains(line, keyword)
	}
}

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	filter := makeFilter("ERROR") // this closure "remembers" the keyword "ERROR"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if filter(line) {
			fmt.Println(line)
		}
	}
}

// it is very common to use closure in a place where you can only pass a certain number of variable but you wanna access few other variables
// just wrap it up in a outer function and the needed function as closure and use the aditional variable
// eg: it is used in gin handler if we need anything other than gin ctx
