package dfa

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cjcodell1/tint/machine"
)

type dfa struct {
	trans   []transition
	start   string
	accepts []string
}

func MakeDFA(trans [][]string, start string, accepts []string) (machine.Machine, error) {
	transitions := []transition{}
	for _, tran := range trans {
		t, err := makeTransition(tran)
		if err != nil {
			return nil, err
		}
		transitions = append(transitions, t)
	}

	return dfa{transitions, start, accepts}, nil
}

func (d dfa) Start(input string) machine.Configuration {
	return config{d.start, strings.Fields(input)}
}

func (d dfa) Step(conf machine.Configuration) (machine.Configuration, error) {
	important, err := conf.GetNext()
	if err != nil {
		return nil, err
	}
	if len(important) != 2 {
		return nil, errors.New("Illegal Configuration")
	}

	// get the current state and symbol
	state := important[0]
	symbol := important[1]

	next_state, err := d.findTransition(state, symbol)
	if err != nil {
		return nil, err
	}

	next_conf, err := conf.Next([]string{next_state})
	if err != nil {
		return nil, err
	}

	return next_conf, nil
}

func (d dfa) IsAccept(conf machine.Configuration) bool {
	if !conf.CanNext() {
		for _, state := range d.accepts {
			if conf.IsState(state) {
				return true
			}
		}
		return false
	}
	return false
}

func (d dfa) IsReject(conf machine.Configuration) bool {
	if !conf.CanNext() {
		for _, state := range d.accepts {
			if conf.IsState(state) {
				return false
			}
		}
		return true
	}
	return false
}

func (d dfa) findTransition(state string, symbol string) (string, error) {
	for _, trans := range d.trans {
		ans, err := trans.IsInput([]string{state, symbol})
		if err != nil {
			return "", err
		}
		if ans {
			output := trans.GetOutput()
			return output[0], nil
		}
	}
	// no transition found
	return "", fmt.Errorf("No transition found for state: \"%s\" and symbol \"%s\"", state, symbol)
}
