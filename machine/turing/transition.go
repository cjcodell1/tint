// Package turing is the implementation of a Turing machine.
package turing

const (
	Left  string = "L"
	Right string = "R"
)

// Transition represents a transition function.
type Transition struct {
	In  Input
	Out Output
}

// Input represents an input to a transition function.
type Input struct {
	State  string
	Symbol string
}

// Output represents an input to a transiiton function.
type Output struct {
	State  string
	Symbol string
	Move   string
}
