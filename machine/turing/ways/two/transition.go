package two

import (
	"errors"
)

const (
	Left  string = "L"
	Right string = "R"
)

// Transition represents a transition function.
type transition struct {
	in  input
	out output
}

// Input represents an input to a transition function.
type input struct {
	state  string
	symbol string
}

// Output represents an input to a transiiton function.
type output struct {
	state  string
	symbol string
	move   string
}

// Output: [state, symbol]
func (t transition) GetInput() []string {
	return []string{t.in.state, t.in.symbol}
}

// Output: [state, symbol, move]
func (t transition) GetOutput() []string {
	return []string{t.out.state, t.out.symbol, t.out.move}
}

// Input: [state, symbol]
func (t transition) IsInput(inputs []string) (bool, error) {
	if len(inputs) != 2 {
		return false, errors.New("Illegal Transition.")
	}
	return t.in.state == inputs[0] && t.in.symbol == inputs[1]
}

// Input: [state, symbol, move]
func (t transition) IsOutput(inputs []string) (bool, error) {
	if len(inputs) != 3 {
		return false, errors.New("Illegal Transition.")
	}
	return t.out.state == inputs[0] && t.out.symbol == inputs[1] && t.out.symbol == inputs[2]
}
