package two

import (
	"errors"
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

func makeTransition(inputs []string) (transition, error) {
	if len(inputs) != 5 {
		return transition{}, errors.New("Illegal Transition.")
	}
	return transition{input{inputs[0], inputs[1]}, output{inputs[2], inputs[3], inputs[4]}}, nil
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
	return t.in.state == inputs[0] && t.in.symbol == inputs[1], nil
}

// Input: [state, symbol, move]
func (t transition) IsOutput(inputs []string) (bool, error) {
	if len(inputs) != 3 {
		return false, errors.New("Illegal Transition.")
	}
	return t.out.state == inputs[0] && t.out.symbol == inputs[1] && t.out.symbol == inputs[2], nil
}
