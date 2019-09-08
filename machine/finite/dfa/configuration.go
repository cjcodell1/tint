package dfa

import (
	"errors"
	"strings"

	"github.com/cjcodell1/tint/machine"
)

type config struct {
	state string
	input []string
}

func (conf config) Print() string {
	var line strings.Builder

	// the WriteString method on a strings.Builder always returns a nil error
	line.WriteString(conf.state)
	line.WriteString(": ")
	line.WriteString(strings.Join(conf.input, " "))
	return line.String()
}

func (conf config) IsState(state string) bool {
	return conf.state == state
}

func (conf config) CanNext() bool {
	return len(conf.input) != 0
}

func (conf config) Next(inputs []string) (machine.Configuration, error) {
	if len(inputs) != 1 {
		return nil, errors.New("Illegal configuration.")
	}

	// Don't step if you can't
	if len(conf.input) == 0 {
		return conf, nil
	}

	// don't want to mutate
	prevInput := make([]string, len(conf.input))
	copy(prevInput, conf.input)

	return config{inputs[0], prevInput[1:]}, nil
}

func (conf config) GetNext() ([]string, error) {
	if len(conf.input) == 0 {
		return nil, errors.New("Illegal Configuration.")
	}
	return []string{conf.state, conf.input[0]}, nil
}
