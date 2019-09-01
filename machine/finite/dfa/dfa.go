package dfa

import (


	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/finite"
)

type dfa struct {
	trans []Transition
	startState string
	acceptStates []string
}

// Makes a DFA. This always returns a nil error. It returns an error to match other MakeXXX funcs.
func MakeDFA(trans []Transition, start string, accepts []string) (machine.Machine, error) {
	return dfa{tras, start, accepts}, nil
}

func (d dfa) Start(input string) machine.Config {
	return finite.Config{d.startState, strings.Fields(input)}, nil
}

func (d dfa) Step(conf machine.Config) (machine.Config, error) {
	finConf, ok := conf.(finite.Config)
	if !ok {
		return nil, errors.New("Connot convert config to correct type for DFAs."}
	}

	// Don't step if you can't
	if len(finConf.String) == 0 {
		return finConf, nil
	}

	// get the current state and symbol
	state := finConf.State
	symbol := finConf.String[0]

	next_state, err := d.findTransition(state, symbol)
	if err != nil {
		return nil, err
	}

	next_conf, err := next(finConf, next_state)
	if err != nil {
		return nil, err
	}

	return next_conf, nil
}

func (d dfa) IsAccept(conf machine.Config) (bool, error) {
	finConf, ok := conf.(finite.Config)
	if !ok {
		return false, errors.New("Cannot convert config to correct type for DFAs.")
	}
	for _, state := range d.acceptStates {
		if finConf.state == state {
			return true, nil
		}
	}
	return false, nil
}

func (d dfa) IsReject(conf machine.Config) (bool, error) {
	accept, err := d.IsAccept(finite.Conf)
	if err != nil {
		return false, err
	}
	return !accept
}

func (d dfa) findTransition(state string, symbol string) (string, error) {
	for _, trans := range d.trans {
		inState, inSymbol := trans.GetInput()
		outState := trans.GetOutput()
		if inState == state {
			if inSymbol == symbol {
				return outState, nil
			}
		}
	}
	// no transition found
	err := fmt.Errorf("No transition found for state: \"%s\" and symbol \"%s\"", state, symbol)
	return "", err
}

func (d dfa) next(conf finite.Config, next_state string) (machine.Config, error) {
	// don't want to mutate conf.String
	prevString := make([]string, len(conf.String))
	copy(prevString, conf.String)

	return Config{next_state, prevString[1:]}
}
