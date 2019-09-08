// Package yaml provides a function to translate a YAML file to a Turing machine.
package yaml

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/cjcodell1/tint/file"
	"github.com/cjcodell1/tint/machine"
)

type builder interface {
	subBuild() (machine.Machine, error)
}

type strawman struct {}
func (s strawman) subBuild() (machine.Machine, error) { return nil, nil }

// Build creates a Turing machine from a YAML file.
func Build(configPath string, machineType string) (machine.Machine, error) {

	var config string
	var err error

	// Read the file
	config, err = file.ReadAll(configPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML
	// TODO figure out how to change the builder in the switch like this
	// 		so I don't have to copy/paste the Unmarshalling and subBuilding each time
	//var b builder

	switch machineType {
		case machine.DFA:
			//b, ok : = b.(dfaBuilder)
			var b dfaBuilder

			err= yaml.Unmarshal([]byte(config), &b)
			if err != nil {
				return nil, err
			}

			// Build the machine
			machine, err := b.subBuild()
			if err != nil {
				return nil, err
			}

			return machine, nil

		case machine.TM:
			//b, ok := b.(tmBuilder)
			var b tmBuilder

			err= yaml.Unmarshal([]byte(config), &b)
			if err != nil {
				return nil, err
			}

			// Build the machine
			machine, err := b.subBuild()
			if err != nil {
				return nil, err
			}

			return machine, nil

		default:
			err = fmt.Errorf("%s is not a valid machine type.", machineType)
			return nil, err
	}
}
