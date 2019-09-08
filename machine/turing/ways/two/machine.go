package two

import (
	"fmt"
	"strings"
	"errors"

	"github.com/cjcodell1/tint/machine"
)


type turingMachine struct {
	trans       []transition
	start  string
	accept string
	reject string
}

// NewTuringMachine is the constructor for a turingMachine.
// It provides error checking necessary for a Turing machine.
// Errors when the accept and reject states are the same state.
func MakeTuringMachine(trans [][]string, start string, accept string, reject string) (machine.Machine, error) {
	if accept == reject {
		return turingMachine{}, fmt.Errorf("%s cannot be both the accept state and the reject state.", accept)
	}
	transitions := []transition{}
	for _, tran := range trans {
		t, err := makeTransition(tran)
		if err != nil {
			return nil, err
		}
		transitions = append(transitions, t)
	}
	return turingMachine{transitions, start, accept, reject}, nil
}

// Start builds the first Config given a space-delimited input string.
func (tm turingMachine) Start(input string) machine.Configuration {
	return configuration{tm.start, strings.Fields(input), 0}
}

// Step applies one transition to the given Config.
// Applies no transition if the Config is in an accept or reject state.
// Errors when there is no transition for the Config.
func (tm turingMachine) Step(conf machine.Configuration) (machine.Configuration, error) {

	// if the state is accept or reject, then don't do anything
	if (tm.IsAccept(conf) || tm.IsReject(conf)) {
		return conf, nil
	}

	next, err := conf.GetNext()
	if err != nil {
		return configuration{}, err
	}
	if len(next) != 2 {
		return configuration{}, errors.New("Illegal configuration.")
	}
	state := next[0]
	symbol := next[1]

	next_state, next_symbol, next_move, err := tm.findTransition(state, symbol)
	if err != nil {
		return nil, err
	}

	next_conf, err := conf.Next([]string{next_state, next_symbol, next_move})
	if err != nil {
		return nil, err
	}

	return next_conf, nil
}

// IsAccept returns true if the Config is in an accept state.
func (tm turingMachine) IsAccept(conf machine.Configuration) bool {
	twoWay, ok := conf.(configuration)
	if !ok {
		return false
	}
	return tm.accept == twoWay.state
}

// IsReject returns true if the Config is in a reject state.
func (tm turingMachine) IsReject(conf machine.Configuration) bool {
	return conf.IsState(tm.reject)
}

func (tm turingMachine) findTransition(state string, symbol string) (string, string, string, error) {
	for _, trans := range tm.trans {
		in := trans.GetInput()
		if len(in) != 2 {
			return "", "", "", errors.New("Illegal transition.")
		}
		inState, inSymbol := in[0], in[1]
		out := trans.GetOutput()
		if len(out) != 3 {
			return "", "", "", errors.New("Illegal transition.")
		}
		outState, outSymbol, outMove := out[0], out[1], out[2]
		if (inState == state) || (inState == machine.Wildcard) {
			if (inSymbol == symbol) || (inSymbol == machine.Wildcard) {
				var next_symbol string
				var next_state string
				if outSymbol == machine.Wildcard {
					next_symbol = symbol // if the output symbol is a wildcard, then re-write the symbol that is on the tape
				} else {
					next_symbol = outSymbol
				}
				if outState == machine.Wildcard {
					next_state = state
				} else {
					next_state = outState
				}
				return next_state, next_symbol, outMove, nil
			}
		}
	}
	// no transition found
	err := fmt.Errorf("no transition found for state: \"%s\" and symbol: \"%s\"", state, symbol)
	return "", "", "", err
}
