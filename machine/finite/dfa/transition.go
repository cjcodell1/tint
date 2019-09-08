package dfa

import (
	"errors"
)

type transition struct {
	in input
	out output
}

type input struct {
	state string
	symbol string
}

type output struct {
	state string
}

func makeTransition(inputs []string) (transition, error) {
	if len(inputs) != 3 {
		return transition{}, errors.New("Illegal Transition.")
	}
	return transition{input{inputs[0], inputs[1]}, output{inputs[2]}}, nil
}

// output: [state, symbol]
func (t transition) GetInput() []string {
	return []string{t.in.state, t.in.symbol}
}

// output: [state]
func (t transition) GetOutput() []string {
	return []string{t.out.state}
}

// input: [state, symbol]
func (t transition) IsInput(inputs []string) (bool, error) {
	if len(inputs) != 2 {
		return false, errors.New("Illegal Transition.")
	}

	return (t.in.state == inputs[0] && t.in.symbol == inputs[1]), nil
}

// input: [state]
func (t transition) IsOutput(inputs []string) (bool, error) {
	if len(inputs) != 1 {
		return false, errors.New("Illegal Transition.")
	}

	return (t.out.state == inputs[0]), nil
}
