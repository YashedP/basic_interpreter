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

func isInteger(str string) bool {

	_, ok := strconv.Atoi(str)
	return ok == nil

}

// LET X = <ARITHMETIC_STATEMENT>
// <ARITHMETIC_STATEMENT> is one of the following: X, X + Y, X - Y, X * Y, or X / Y
func basicLet(line statement) {

	var letter string
	letter = string(line.args[0])
	var strArray []string
	strArray = strings.Fields(line.args)
	var firstTerm string
	var secondTerm string
	var intFirstTerm int
	var intSecondTerm int

	// This if statement checks if there are any operations and if not, stores the value in variable
	if len(strArray) < 4 {
		if isInteger(strArray[2]) {
			variables[letter], _ = strconv.Atoi(strArray[2])
		} else {
			variables[letter] = variables[strArray[2]]
		}
		return
	}

	firstTerm = strArray[2]
	secondTerm = strArray[4]
	intFirstTerm = 0
	intSecondTerm = 0

	// Processing Operations
	if strings.Contains(line.args, "+") {
		if isInteger(firstTerm) && isInteger(secondTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			variables[letter] = intFirstTerm + intSecondTerm

		} else if isInteger(firstTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm = variables[secondTerm]
			variables[letter] = intFirstTerm + intSecondTerm

		} else if isInteger(secondTerm) {
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm + intSecondTerm

		} else {
			intSecondTerm = variables[secondTerm]
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm + intSecondTerm
		}

	} else if strings.Contains(line.args, "-") {
		if isInteger(firstTerm) && isInteger(secondTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			variables[letter] = intFirstTerm - intSecondTerm

		} else if isInteger(firstTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm = variables[secondTerm]
			variables[letter] = intFirstTerm - intSecondTerm

		} else if isInteger(secondTerm) {
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm - intSecondTerm

		} else {
			intSecondTerm = variables[secondTerm]
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm - intSecondTerm
		}

	} else if strings.Contains(line.args, "*") {
		if isInteger(firstTerm) && isInteger(secondTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			variables[letter] = intFirstTerm * intSecondTerm

		} else if isInteger(firstTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm = variables[secondTerm]
			variables[letter] = intFirstTerm * intSecondTerm

		} else if isInteger(secondTerm) {
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm * intSecondTerm

		} else {
			intSecondTerm = variables[secondTerm]
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm * intSecondTerm
		}

	} else if strings.Contains(line.args, "/") {
		if isInteger(firstTerm) && isInteger(secondTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			variables[letter] = intFirstTerm / intSecondTerm

		} else if isInteger(firstTerm) {
			intFirstTerm, _ = strconv.Atoi(firstTerm)
			intSecondTerm = variables[secondTerm]
			variables[letter] = intFirstTerm / intSecondTerm

		} else if isInteger(secondTerm) {
			intSecondTerm, _ = strconv.Atoi(secondTerm)
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm / intSecondTerm

		} else {
			intSecondTerm = variables[secondTerm]
			intFirstTerm = variables[firstTerm]
			variables[letter] = intFirstTerm / intSecondTerm
		}
	}
}

// IF <CONDITION> THEN GOTO L
// <CONDITION> is one of the following: X = Y, X > Y, X < Y, X <> Y, X <= Y, or X >= Y
func basicIf(line statement, index int) int {
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
	} else if val, ok := undefinedVariables[line.args]; ok {
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
	} else if val, ok := variables[line.args]; ok { // condition for printing a variable that exists
		println(val)
	} else if val, ok := undefinedVariables[line.args]; ok {
		println(val)
	} else
}

func main() {
	// Initialize the variables hashmap
	variables = make(map[string]int)

	read_lines()
	sort_lines()

	for i := 0; i < len(code); i++ {
		println(code[i].label, code[i].command, code[i].args)
	}

	// for i := 0; i < len(code); i++ {
	// 	index := interpret(i)
	// 	if index != 0 {
	// 		i = index
	// 	}
	// }
}
