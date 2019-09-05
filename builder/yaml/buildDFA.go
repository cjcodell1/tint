package yaml

import (
	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/finite/dfa"
)

// dfaBuilder is the struct to marshal the YAML.
type dfaBuilder struct {
	// These must be export, yaml parser requires it.
	Start string
	Accepts []string `yaml:"accept-states"` // renamed to accept-states
	Transitions [][]string
}

func (b dfaBuilder) subBuild() (machine.Machine, error) {
	d, err := dfa.MakeDFA(b.Transitions, b.Start, b.Accepts)
	if err != nil {
		return nil, err
	}

	return d, nil
}
