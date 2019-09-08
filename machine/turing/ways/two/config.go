package two

import (
	"fmt"
	"strings"
	"errors"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
)

type configuration struct {
	state string
	tape  []string
	head int
}

func (conf configuration) Print() string {
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


func (conf configuration) IsState(state string) bool {
	return conf.state == state
}

func (conf configuration) CanNext() bool {
	return true
}

func (conf configuration) Next(inputs []string) (machine.Configuration, error) {
	if len(inputs) != 3 {
		return nil, errors.New("Illegal configuration.")
	}

	next_state := inputs[0]
	next_symbol := inputs[1]
	next_move := inputs[2]

	// Assume that conf is not in an accept or a reject state.

	// Don't want to mutate
	prevTape := make([]string, len(conf.tape))
	copy(prevTape, conf.tape)

	var next_conf configuration
	// transition to the next state
	next_conf.state = next_state

	leng := len(prevTape)
	head := conf.head

	// write the next symbol
	if (head == leng) && (next_symbol == turing.Blank) {
		next_conf.tape = prevTape
	} else if (head == leng) && (next_symbol != turing.Blank) {
		next_conf.tape = append(prevTape, next_symbol)
	} else if (head == 0) && (next_symbol == turing.Blank) {
		// replace first symbol with a Blank and DO NOT move
		next_conf.tape = append(prevTape[:0], prevTape[1:]...)
		next_conf.head = head
		return next_conf, nil
	} else {
		next_conf.tape = prevTape
		next_conf.tape[conf.head] = next_symbol
	}

	// move in the next direction

	// move AFTER write, so take another len with the next tape
	leng = len(next_conf.tape)

	if (head == leng) && (next_move == turing.Right) {
		next_conf.head = head
	} else if (head == 0) && (next_move == turing.Left) {
		next_conf.head = head
	} else {
		if next_move == turing.Right {
			next_conf.head = head + 1
		} else if next_move == turing.Left {
			next_conf.head = head - 1
		} else {
			return configuration{}, fmt.Errorf("%s is not a legal move, use %s or %s", next_move, turing.Right, turing.Left)
		}
	}

	return next_conf, nil
}

func (conf configuration) GetNext() ([]string, error) {
	if conf.head < len(conf.tape) {
		return []string{conf.state, conf.tape[conf.head]}, nil
	}
	return []string{conf.state, turing.Blank}, nil
}
