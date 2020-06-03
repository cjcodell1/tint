package yaml

import (
	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing/ways/two"
	"github.com/cjcodell1/tint/machine/turing/ways/one"
)

type oneWayTmBuilder struct {
	// These must be exported, yaml parser requires it.
	Start       string
	Accept      string
	Reject      string
	Transitions [][]string
}

func (b oneWayTmBuilder) subBuild() (machine.Machine, error) {
	tm, err := one.MakeTuringMachine(b.Transitions, b.Start, b.Accept, b.Reject)
	if err != nil {
		return nil, err
	}

	return tm, nil
}

type twoWayTmBuilder struct {
	// These must be exported, yaml parser requires it.
	Start       string
	Accept      string
	Reject      string
	Transitions [][]string
}

func (b twoWayTmBuilder) subBuild() (machine.Machine, error) {
	tm, err := two.MakeTuringMachine(b.Transitions, b.Start, b.Accept, b.Reject)
	if err != nil {
		return nil, err
	}

	return tm, nil
}
