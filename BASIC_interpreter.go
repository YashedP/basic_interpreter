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
		basicLet(line)
	case "IF":
		index = basicIf(line, index)
	case "PRINT":
		basicPrint(line)
	case "PRINTLN":
		basicPrintln(line)
	}

	return index
}

// Implementation of the LET command
func basicLet(line statement) {

}

// Implementation of the IF command
func basicIf(line statement, index int) int {
	return index
}

// Implementation of the PRINT command
func basicPrint(line statement) {

}

// Implementation of the PRINTLN command
func basicPrintln(line statement) {

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
