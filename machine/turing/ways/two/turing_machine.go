package two

import (
	"fmt"
	"strings"
	"errors"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
	"github.com/cjcodell1/tint/machine/turing/ways"
)


type turingMachine struct {
	trans       []ways.Transition
	startState  string
	acceptState string
	rejectState string
}

// NewTuringMachine is the constructor for a turingMachine.
// It provides error checking necessary for a Turing machine.
// Errors when the accept and reject states are the same state.
func MakeTuringMachine(trans []ways.Transition, start string, accept string, reject string) (machine.Machine, error) {
	if accept == reject {
		return turingMachine{}, fmt.Errorf("%s cannot be both the accept state and the reject state.", accept)
	}
	return turingMachine{trans, start, accept, reject}, nil
}

// Start builds the first Config given a space-delimited input string.
func (tm turingMachine) Start(input string) machine.Config {
	return twoWayConfig{tm.startState, strings.Fields(input), 0}
}

// Step applies one transition to the given Config.
// Applies no transition if the Config is in an accept or reject state.
// Errors when there is no transition for the Config.
func (tm turingMachine) Step(conf machine.Config) (machine.Config, error) {
	twoWay, ok := conf.(twoWayConfig)
	if !ok {
		errors.New("Cannot convert config to correct type for two-way Turing machines.")
	}
	state := twoWay.state

	// if the state is accept or reject, then don't do anything
	accept, err := tm.IsAccept(conf)
	if err != nil {
		return nil, err
	}
	reject, err := tm.IsReject(conf)
	if err != nil {
		return nil, err
	}

	if (accept || reject) {
		return conf, nil
	}

	var symbol string
	if twoWay.head == len(twoWay.tape) {
		symbol = turing.Blank
	} else {
		symbol = twoWay.tape[twoWay.head]
	}

	next_state, next_symbol, next_move, err := tm.findTransition(state, symbol)
	if err != nil {
		fmt.Printf("STEP: %v\n", conf)
		fmt.Printf("state: %s\nsymbol: %s\n", state, symbol)
		fmt.Println("ERROR2")
		return twoWayConfig{}, err
	}

	next_conf, err := next(twoWay, next_state, next_symbol, next_move)
	if err != nil {
		return twoWayConfig{}, err
	}

	return next_conf, nil
}

// IsAccept returns true if the Config is in an accept state.
func (tm turingMachine) IsAccept(conf machine.Config) (bool, error) {
	twoWay, ok := conf.(twoWayConfig)
	if !ok {
		return false, errors.New("Cannot convert config to correct type for two-way Turing machines.")
	}
	return tm.acceptState == twoWay.state, nil
}

// IsReject returns true if the Config is in a reject state.
func (tm turingMachine) IsReject(conf machine.Config) (bool, error) {
	twoWay, ok := conf.(twoWayConfig)
	if !ok {
		return false, errors.New("Cannot convert config to correct type for two-way Turing machines.")
	}
	return tm.rejectState == twoWay.state, nil
}

func (tm turingMachine) findTransition(state string, symbol string) (string, string, string, error) {
	for _, trans := range tm.trans {
		inState, inSymbol := trans.GetInput()
		outState, outSymbol, outMove := trans.GetOutput()
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

func next(conf twoWayConfig, next_state string, next_symbol string, next_move string) (machine.Config, error) {

	// don't want to mutate conf.Tape
	var prevTape = make([]string, len(conf.tape))
	copy(prevTape, conf.tape)

	var next_conf twoWayConfig
	// transition to the next state
	next_conf.state = next_state

	leng := len(prevTape)
	head := conf.head

	// write the next symbol
	if (head == leng) && (next_symbol == turing.Blank) {
		next_conf.tape = prevTape
	} else if (head == leng) && (next_symbol != turing.Blank) {
		next_conf.tape = append(prevTape, next_symbol)
	} else if (head == 0) && (next_symbol == turing.Blank) {
		// replace first symbol with a Blank and DO NOT move
		next_conf.tape = append(prevTape[:0], prevTape[1:]...)
		next_conf.head = head
		return next_conf, nil
	} else {
		next_conf.tape = prevTape
		next_conf.tape[conf.head] = next_symbol
	}

	// move in the next direction

	// move AFTER write, so take another len with the next tape
	leng = len(next_conf.tape)

	if (head == leng) && (next_move == ways.Right) {
		next_conf.head = head
	} else if (head == 0) && (next_move == ways.Left) {
		next_conf.head = head
	} else {
		if next_move == ways.Right {
			next_conf.head = head + 1
		} else if next_move == ways.Left {
			next_conf.head = head - 1
		} else {
			return twoWayConfig{}, fmt.Errorf("%s is not a legal move, use %s or %s", next_move, ways.Right, ways.Left)
		}
	}

	return next_conf, nil
}
