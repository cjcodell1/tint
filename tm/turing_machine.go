package tm

import (
	"fmt"
	"strings"
	"log"
)

const (
	Blank string = "_"
)


type TuringMachine interface {
	Start(input string) Config
	Step(conf Config) Config
}

type turingMachine struct {
	Trans []Transition
	StartState string
	AcceptState string
	RejectState string
}

func NewTuringMachine(trans []Transition, start string, accept string, reject string) turingMachine {
	if accept == reject {
		log.Fatal("Accept state and reject state cannot be the same state.")
	}
	return turingMachine{trans, start, accept, reject}
}

func (tm turingMachine) Start(input string) Config {
	return Config{tm.StartState, strings.Fields(input), 0}
}

func (tm turingMachine) Step(conf Config) Config {
	state := conf.State

	// if the state is accept or reject, then don't do anything
	if (state == tm.AcceptState) || (state == tm.RejectState) {
		return conf
	}

	var symbol string
	if conf.Index == len(conf.Tape) {
		symbol = Blank
	} else {
		symbol = conf.Tape[conf.Index]
	}

	next_state, next_symbol, next_move, err := tm.findTransition(state, symbol)

	if err != nil {
		panic(err)
	}
	return next(conf, next_state, next_symbol, next_move)
}

func (tm turingMachine) findTransition(state string, symbol string) (string, string, string, error) {
	for _, trans := range tm.Trans {
		if (trans.In == Input{state, symbol}) {
			return trans.Out.State, trans.Out.Symbol, trans.Out.Move, nil
		}
	}
	// no transition found
	err := fmt.Errorf("no transition found for state: \"%s\" and symbol: \"%s\"", state, symbol)
	return "", "", "", err
}

func next(conf Config, next_state string, next_symbol string, next_move string) Config {
	var next_conf Config
	// transition to the next state
	next_conf.State = next_state

	leng := len(conf.Tape)
	index := conf.Index

	// write the next symbol
	if (index == leng) && (next_symbol == Blank) {
		next_conf.Tape = conf.Tape
	} else if (index == leng) && (next_symbol != Blank) {
		next_conf.Tape = append(conf.Tape, next_symbol)
	} else if (index == 0) && (next_symbol == Blank) {
		// replace first symbol with a Blank and DO NOT move
		next_conf.Tape = append(conf.Tape[:0], conf.Tape[1:] ...)
		next_conf.Index = index
		return next_conf
	} else {
		next_conf.Tape = conf.Tape
		next_conf.Tape[conf.Index] = next_symbol
	}

	// move in the next direction

	// move AFTER write, so take another len with the next tape
	leng = len(next_conf.Tape)

	if (index == leng) && (next_move == Right) {
		next_conf.Index = index
	} else if (index == 0) && (next_move == Left) {
		next_conf.Index = index
	} else {
		if next_move == Right {
			next_conf.Index = index + 1
		} else {
			next_conf.Index = index - 1
		}
	}

	return next_conf
}
