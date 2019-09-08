package yaml

import (
	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing/ways/two"
)

// turingMachine is the struct to use to marshal the YAML.
type tmBuilder struct {
	// These must be exported, yaml parser requires it.
	Start       string
	Accept      string
	Reject      string
	Transitions [][]string
}

func (b tmBuilder) subBuild() (machine.Machine, error) {
	tm, err := two.MakeTuringMachine(b.Transitions, b.Start, b.Accept, b.Reject)
	if err != nil {
		return nil, err
	}

	return tm, nil
}
