package dfa

import (
	"errors"

	"github.com/cjcodell1/tint/machine"
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

// output: [state, symbol]
func (t machine.Transition) GetInput() []string {
	return []string{t.in.state, t.in.symbol}
}

// output: [state]
func (t machine.Transition) GetOutput() []string {
	return []string{t.out.state}
}

// input: [state, symbol]
func (t machine.Transition) IsInput(inputs []string) (bool, error) {
	if len(inputs) != 2 {
		return false, errors.New("Illegal Transition.")
	}

	return (t.in.state == inputs[0] && t.in.symbol == inputs[1]), nil
}

// input: [state]
func (t machine.Transition) IsOutput(inputs []string) (bool, error) {
	if len(inputs) != 1 {
		return false, errors.New("Illegal Transition.")
	}

	return (t.out.state == inputs[0]), nil
}
