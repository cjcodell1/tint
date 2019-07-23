package one

import (
	"strings"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
)

type oneWayConfig struct {
	state string
	tape  []string
	head int
}

func (conf waysConfig) Print() string {
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

	// add the tape with a blank at the end
	// add spaces for tape up until where the head is supposed to go
	carrot := 0
	for _, sym := range conf.tape {
		line1.WriteString(" ")
		line1.WriteString(sym)

		if carrot == conf.head {
			line2.WriteString("^")
		} else if carrot < conf.head {
			carrot++
			line2.WriteString(" ")
			for _, _ = range sym {
				line2.WriteString(" ")
			}
		}
	}
	line1.WriteString(" ")
	line1.WriteString(turing.Blank)

	return line1.String() + "\n" + line2.String()
}
