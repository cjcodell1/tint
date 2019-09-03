package dfa

import (


	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/finite"
)

type dfa struct {
	trans []transition
	start string
	accepts []string
}

func MakeDFA(trans []string, start string, accepts []string) (machine.Machine, error) {
	transitions := []transition{}
	for _, tran := range trans {
		t, err := makeTransition(tran)
		if err != nil {
			return nil, err
		}
		transitions = append(transitions, makeTransition(t)
	}

	return dfa{transitions, start, accepts}, nil
}

func (d dfa) Start(input string) machine.Config {
	return finite.Config{d.start, strings.Fields(input)}, nil
}

func (d dfa) Step(conf machine.Config) (machine.Config, error) {
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

	next_conf, err := conf.Next(next_state)
	if err != nil {
		return nil, err
	}

	return next_conf, nil
}

func (d dfa) IsAccept(conf machine.Config) bool {
	if !conf.CanNext() {
		for _, state := range d.accepts {
			if conf.IsState(state) {
				return true
			}
		}
	}
	return false
}

func (d dfa) IsReject(conf machine.Config) bool {
	if !conf.CanNext() {
		for _, state := range d.accepts {
			if conf.IsState(state) {
				return false
			}
		}
	}
	return true
}

func makeTransition(inputs []string) {
	if len(inputs) != 3 {
		return nil, errors.New("Illegal Transition.")
	}
	return transition{input{inputs[0], inputs[1]}, output{inputs[2]}}
}

func (d dfa) findTransition(state string, symbol string) (string, error) {
	for _, trans := range d.trans {
		ans, err := trans.IsInput([]string{state, symbol})
		if err != nil {
			return "", err
		}
		output := trans.GetOutput()
		return output[0], nil
	}
	// no transition found
	return "", fmt.Errorf("No transition found for state: \"%s\" and symbol \"%s\"", state, symbol)
}
