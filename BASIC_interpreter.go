package main

import (
	"bufio"
	"os"
	"strings"
)

type statement struct {
	label   string
	command string
	args    string
}

var (
	code      []statement
	variables map[string]int
)

// Read the next line from the input and store it in the code variable
// If the input stream is empty, return an empty statement
// If the input stream is not empty, return the next statement
func read_lines() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			break
		}

		if err := scanner.Err(); err != nil {
			break
		}

		arr := strings.SplitN(line, " ", 3)

		code = append(code, statement{
			label:   arr[0],
			command: arr[1],
			args:    arr[2],
		})
	}
}

// Calls the right function to interpret the statement
func interpret(index int) int {
	line := code[index]

	index = 0

	switch line.command {
	case "LET":
		basicLet()
	case "IF":
		index = basicIf(index)
	case "PRINT":
		basicPrint()
	case "PRINTLN":
		basicPrintln()
	}

	return index
}

// Implementation of the LET command
func basicLet() {

}

// Implementation of the IF command
func basicIf(index int) int {
	return index
}

// Implementation of the PRINT command
func basicPrint() {

}

// Implementation of the PRINTLN command
func basicPrintln() {

}

func main() {
	// Initialize the variables hashmap
	variables = make(map[string]int)

	read_lines()

	for i := 0; i < len(code); i++ {
		index := interpret(i)
		if index != 0 {
			i = index
		}
	}

	// // If the line is not empty, interpret it
	// if line.label != "" {
	// 	interpret(line)
	// }
}
