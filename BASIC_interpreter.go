package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type statement struct {
	label   int
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
			label:   func() int { i, _ := strconv.Atoi(arr[0]); return i }(),
			command: arr[1],
			args:    arr[2],
		})
	}
}

func sort_lines() {
	// Sort the code array by label
	sort.Slice(code, func(i, j int) bool {
		return code[i].label < code[j].label
	})
}

// Calls the right function to interpret the statement
func interpret(index int) int {
	line := code[index]
	index = index + 1

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

func isInteger(str string) bool {
	_, ok := strconv.Atoi(str)
	return ok == nil

}

// LET X = <ARITHMETIC_STATEMENT>
// <ARITHMETIC_STATEMENT> is one of the following: X, X + Y, X - Y, X * Y, or X / Y
func basicLet(line statement) {
	letter := string(line.args[0])
	strArray := strings.Fields(line.args)
	var firstTerm int
	var secondTerm int

	// This if statement checks if there are any operations and if not, stores the value in variable
	if len(strArray) < 4 {
		if isInteger(strArray[2]) {
			variables[letter], _ = strconv.Atoi(strArray[2])
		} else {
			variables[letter] = variables[strArray[2]]
		}
		return
	}

	// Converts the argument to integer values
	if isInteger(strArray[2]) {
		firstTerm, _ = strconv.Atoi(strArray[2])
	} else {
		firstTerm = variables[strArray[2]]
	}

	if isInteger(strArray[4]) {
		secondTerm, _ = strconv.Atoi(strArray[4])
	} else {
		secondTerm = variables[strArray[4]]
	}

	// Processing Operations
	if strings.Contains(line.args, "+") {
		variables[letter] = firstTerm + secondTerm

	} else if strings.Contains(line.args, "-") {
		variables[letter] = firstTerm - secondTerm

	} else if strings.Contains(line.args, "*") {
		variables[letter] = firstTerm * secondTerm

	} else if strings.Contains(line.args, "/") {
		variables[letter] = firstTerm / secondTerm

	}
}

// IF <CONDITION> THEN GOTO L
// <CONDITION> is one of the following: X = Y, X > Y, X < Y, X <> Y, X <= Y, or X >= Y
func basicIf(line statement, index int) int {

	strArray := strings.Fields(line.args)
	var firstTerm int
	var secondTerm int
	var operator string

	if isInteger(strArray[0]) {
		firstTerm, _ = strconv.Atoi(strArray[0])
	} else {
		firstTerm = variables[strArray[0]]
	}
	if isInteger(strArray[2]) {
		secondTerm, _ = strconv.Atoi(strArray[2])
	} else {
		secondTerm = variables[strArray[2]]
	}
	operator = strArray[1]

	switch operator {
	case "=":
		if firstTerm == secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
			index = index - 1
		}

	case ">":
		if firstTerm > secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
			index = index - 1
		}

	case "<":
		if firstTerm < secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
			index = index - 1
		}

	case "<>":
		if firstTerm == secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
			index = index - 1
		}

	case "<=":
		if firstTerm <= secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
			index = index - 1
		}

	case ">=":
		if firstTerm >= secondTerm {
			index, _ = strconv.Atoi(strArray[5])
			index = index / 10
		}

	}

	return index
}

// PRINT <PRINT_STATEMENT>
// <PRINT_STATEMENT> is either a variable name or a literal string delimited by double quotes
// Inside the quotes, the string contains only alphanumeric characters (a-z, A-Z, 0-9) and spaces
func basicPrint(line statement) {
	// condition for printing a string; the reason it's compared to 34 is because str[index] returns the byte value
	if len(line.args) > 0 && line.args[0] == 34 && line.args[len(line.args)-1] == 34 {
		print(line.args[1 : len(line.args)-1])
	} else if val, ok := variables[line.args]; ok { // condition for printing a variable that exists
		print(val)
	}
}

// PRINTLN <PRINT_STATEMENT>
// <PRINT_STATEMENT> is either a variable name or a literal string delimited by double quotes
// Inside the quotes, the string contains only alphanumeric characters (a-z, A-Z, 0-9) and spaces
func basicPrintln(line statement) {
	// condition for printing a string; the reason it's compared to 34 is because str[index] returns the byte value
	if len(line.args) > 0 && line.args[0] == 34 && line.args[len(line.args)-1] == 34 {
		println(line.args[1 : len(line.args)-1])
	} else if val, ok := variables[line.args]; ok {
		println(val)
	}
}

func main() {
	// Initialize the variables hashmap
	variables = make(map[string]int)

	read_lines()
	sort_lines()

	for i := 0; i < len(code); {
		i = interpret(i)

	}
}
