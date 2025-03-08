package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type statement struct {
	label   int
	command string
	args    []string
}

var (
	code      []statement
	variables map[string]int
	labels    map[int]int
)

// Read the next line from the input and store it in the code variable
// If the input stream is empty, return an empty statement
// If the input stream is not empty, return the next statement
func read_lines() {
	scanner := bufio.NewScanner(os.Stdin)

	// file, _ := os.Open("test1")
	// scanner := bufio.NewScanner(file)

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

		str := arr[2]

		var args = []string{}

		switch arr[1] {
		case "LET":
			args = strings.Fields(str)
		case "IF":
			args = strings.Fields(str)
		default:
			args = []string{str}
		}

		code = append(code, statement{
			label:   func() int { i, _ := strconv.Atoi(arr[0]); return i }(),
			command: arr[1],
			args:    args,
		})
	}
}

func add_labels() {
	// Add the labels to the code array
	for i := 0; i < len(code); i++ {
		labels[code[i].label] = i
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

	switch line.command[0] {
	case 'L': // LET
		if line.command == "LET" {
			basicLet(line)
		}
	case 'I': // IF
		if line.command == "IF" {
			index = basicIf(line, index)
		}
	case 'P': // PRINT or PRINTLN
		switch line.command[len(line.command)-1] {
		case 'T':
			basicPrint(line)
		case 'N':
			basicPrintln(line)
		}
	}

	return index
}

func isInteger(str string) bool {
	// Check using ascii values of num range
	if str[0] >= 48 && str[0] <= 57 {
		return true
	}
	return false
}

// LET X = <ARITHMETIC_STATEMENT>
// <ARITHMETIC_STATEMENT> is one of the following: X, X + Y, X - Y, X * Y, or X / Y
func basicLet(line statement) {
	switch len(line.args) {
	case 1:
		if isInteger(line.args[0]) {
			variables[line.args[0]], _ = strconv.Atoi(line.args[2])
		} else {
			variables[line.args[0]] = variables[line.args[2]]
		}
	case 3:
		variables[line.args[0]] = func() int { i, _ := strconv.Atoi(line.args[2]); return i }()
	case 5:
		var firstTerm, secondTerm int

		// Convert first term
		if isInteger(line.args[2]) {
			firstTerm, _ = strconv.Atoi(line.args[2])
		} else {
			firstTerm = variables[line.args[2]]
		}

		// Convert second term
		if isInteger(line.args[4]) {
			secondTerm, _ = strconv.Atoi(line.args[4])
		} else {
			secondTerm = variables[line.args[4]]
		}

		// Perform the operation based on the operator
		switch line.args[3] {
		case "+":
			variables[line.args[0]] = firstTerm + secondTerm
		case "-":
			variables[line.args[0]] = firstTerm - secondTerm
		case "*":
			variables[line.args[0]] = firstTerm * secondTerm
		case "/":
			if secondTerm != 0 {
				variables[line.args[0]] = firstTerm / secondTerm
			}
		}
	}
}

// IF <CONDITION> THEN GOTO L
// <CONDITION> is one of the following: X = Y, X > Y, X < Y, X <> Y, X <= Y, or X >= Y
func basicIf(line statement, index int) int {
	var firstTerm, secondTerm int

	if isInteger(line.args[0]) {
		firstTerm, _ = strconv.Atoi(line.args[0])
	} else {
		firstTerm = variables[line.args[0]]
	}
	if isInteger(line.args[2]) {
		secondTerm, _ = strconv.Atoi(line.args[2])
	} else {
		secondTerm = variables[line.args[2]]
	}

	condition := false
	switch line.args[1] {
	case "=":
		if firstTerm == secondTerm {
			condition = true
		}
	case ">":
		if firstTerm > secondTerm {
			condition = true
		}
	case "<":
		if firstTerm < secondTerm {
			condition = true
		}
	case "<>":
		if firstTerm == secondTerm {
			condition = true
		}
	case "<=":
		if firstTerm <= secondTerm {
			condition = true
		}
	case ">=":
		if firstTerm >= secondTerm {
			condition = true
		}
	}

	if condition {
		index = labels[func() int { i, _ := strconv.Atoi(line.args[5]); return i }()]
	}

	return index
}

// PRINT <PRINT_STATEMENT>
// <PRINT_STATEMENT> is either a variable name or a literal string delimited by double quotes
// Inside the quotes, the string contains only alphanumeric characters (a-z, A-Z, 0-9) and spaces
func basicPrint(line statement) {
	str := line.args[0]
	// condition for printing a string; the reason it's compared to 34 is because str[index] returns the byte value
	if len(str) > 0 && str[0] == 34 && str[len(str)-1] == 34 {
		fmt.Print(str[1 : len(str)-1])
	} else if val, ok := variables[str]; ok { // condition for printing a variable that exists
		fmt.Print(val)
	}
}

// PRINTLN <PRINT_STATEMENT>
// <PRINT_STATEMENT> is either a variable name or a literal string delimited by double quotes
// Inside the quotes, the string contains only alphanumeric characters (a-z, A-Z, 0-9) and spaces
func basicPrintln(line statement) {
	str := line.args[0]
	// condition for printing a string; the reason it's compared to 34 is because str[index] returns the byte value
	if len(str) > 0 && str[0] == 34 && str[len(str)-1] == 34 {
		fmt.Println(str[1 : len(str)-1])
	} else if val, ok := variables[str]; ok {
		fmt.Println(val)
	}
}

func main() {
	// Initialize the variables hashmap
	variables = make(map[string]int)
	labels = make(map[int]int)

	read_lines()
	sort_lines()
	add_labels()

	for i := 0; i < len(code); {
		i = interpret(i)
	}
}
