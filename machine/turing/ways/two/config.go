package two

import (
	"strings"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
)

type config struct {
	state string
	tape  []string
	head int
}

func (conf config) Print() string {
	var line1 strings.Builder
	var line2 strings.Builder

	// the WriteString method on a strings.Builder always returns a nil error.

	// add state and semicolon
	line1.WriteString(conf.state)
	line1.WriteString(":")

	// add spaces for the state and semicolon
	for _, _ = range conf.state {
		line2.WriteString(" ")
	}
	line2.WriteString(" ")

	// add the tape with a blank at the beginning and end
	// add spaces for tape up until where the head is supposed to go
	line1.WriteString(" ")
	line1.WriteString(turing.Blank)
	line2.WriteString("  ")

	// now write what's on the tape
	line1.WriteString(" ")
	line1.WriteString(strings.Join(conf.tape, " "))

	carrot := 0
	for {
		if carrot == conf.head {
			line2.WriteString("^")
			break
		} else {
			line2.WriteString(" ")
			for _, _ = range conf.tape[carrot] {
				line2.WriteString(" ")
			}
			carrot += 1
		}
	}

	// write the last blank
	line1.WriteString(" ")
	line1.WriteString(turing.Blank)

	return line1.String() + "\n" + line2.String()
}


func (conf config) IsState(state string) bool {
	return conf.state == state
}

func (conf config) CanNext() bool {
	return true
}

func (conf config) Next(inputs []string) (machine.Configuration, error) {

}

func (conf config) GetNext() ([]string, error) {
	if conf.head < len(conf.tape) {
		return []string{conf.state, conf.tape[conf.head]}, nil
	}
	return []string{conf.state, turing.Blank}, nil
}
