package main

type statement struct {
	label   string
	command string
	args    []string
}

var (
	code      []statement
	variables map[string]int
)

// Read the next line from the input and store it in the code variable
// If the input stream is empty, return an empty statement
// If the input stream is not empty, return the next statement
func next_line() statement {
	return statement{}
}

// Calls the right function to interpret the statement
func interpret(statement statement) {
	switch statement.command {
	case "LET":
		basicLet()
	case "IF":
		basicIf()
	case "PRINT":
		basicPrint()
	case "PRINTLN":
		basicPrintln()
	}
}

// Implementation of the LET command
func basicLet() {

}

// Implementation of the IF command
func basicIf() {

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

	// Read the first line
	line := next_line()

	// If the line is not empty, interpret it
	if line.label != "" {
		interpret(line)
	}
}
