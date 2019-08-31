package yaml

import (
	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing/ways"
	"github.com/cjcodell1/tint/machine/turing/ways/two"
)

// turingMachine is the struct to use to marshal the YAML.
type tmBuilder struct {
	// These must be exported, yaml parser requires it.
	Start       string
	Accept      string
	Reject      string
	Transitions [][5]string
}

func (b tmBuilder) subBuild() (machine.Machine, error) {
	var trans []ways.Transition
	for _, t := range b.Transitions {
		trans = append(trans, ways.MakeTransition(t[0], t[1], t[2], t[3], t[4]))
	}
	tm, err := two.MakeTuringMachine(trans, b.Start, b.Accept, b.Reject)
	if err != nil {
		return nil, err
	}

	return tm, nil
}
