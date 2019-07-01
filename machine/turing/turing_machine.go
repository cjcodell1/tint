// Package turing is the implementation of a Turing machine.
package turing

import (
    "fmt"
    "strings"
)

const (
    Blank string = "_"
    Wildcard string = "*"
)


// TuringMachine is the interface specifying which functions are available for Turing Machines.
type TuringMachine interface {
    Start(input string) Config
    Step(conf Config) (Config, error)
    IsAccept(conf Config) bool
    IsReject(conf Config) bool
}

// turingMachine is an UNEXPORTED struct with EXPORTED fields,
// forcing programmers to use the constructor which provides error checking.
type turingMachine struct {
    Trans []Transition
    StartState string
    AcceptState string
    RejectState string
}

// NewTuringMachine is the constructor for a TuringMachine.
// It provides error checking necessary for a Turing machine.
// Errors when the accept and reject states are the same state.
func NewTuringMachine(trans []Transition, start string, accept string, reject string) (TuringMachine, error) {
    if accept == reject {
        return turingMachine{}, fmt.Errorf("%s cannot be both the accept state and the reject state.", accept)
    }
    return turingMachine{trans, start, accept, reject}, nil
}

// Start builds the first Config given an input string.
func (tm turingMachine) Start(input string) Config {
    return Config{tm.StartState, strings.Fields(input), 0}
}

// Step applies one transition to the given Config.
// Applies no transition if the Config is in an accept or reject state.
// Errors when there is no transition for the Config.
func (tm turingMachine) Step(conf Config) (Config, error) {
    state := conf.State

    // if the state is accept or reject, then don't do anything
    if (tm.IsAccept(conf)) || (tm.IsReject(conf)) {
        return conf, nil
    }

    var symbol string
    if conf.Index == len(conf.Tape) {
        symbol = Blank
    } else {
        symbol = conf.Tape[conf.Index]
    }

    next_state, next_symbol, next_move, err := tm.findTransition(state, symbol)
    if err != nil {
        return Config{}, err
    }
    return next(conf, next_state, next_symbol, next_move), nil
}

// IsAccept returns true if the Config is in an accept state.
func (tm turingMachine) IsAccept(conf Config) bool {
    return tm.AcceptState == conf.State
}

// IsReject returns true if the Config is in a reject state.
func (tm turingMachine) IsReject(conf Config) bool {
    return tm.RejectState == conf.State
}


func (tm turingMachine) findTransition(state string, symbol string) (string, string, string, error) {
    for _, trans := range tm.Trans {
        if (trans.In.State == state) || (trans.In.State == Wildcard) {
            if (trans.In.Symbol == symbol) || (trans.In.Symbol == Wildcard) {
                if trans.Out.Symbol == Wildcard {
                    return trans.Out.State, symbol, trans.Out.Move, nil // if the output symbol is a wildcard, then re-write the symbol that is on the tape
                } else {
                    return trans.Out.State, trans.Out.Symbol, trans.Out.Move, nil
                }
            }
        }
    }
    // no transition found
    err := fmt.Errorf("no transition found for state: \"%s\" and symbol: \"%s\"", state, symbol)
    return "", "", "", err
}

func next(conf Config, next_state string, next_symbol string, next_move string) Config {
    // don't want to mutate conf.Tape
    var prevTape = make([]string, len(conf.Tape))
    copy(prevTape, conf.Tape)


    var next_conf Config
    // transition to the next state
    next_conf.State = next_state

    leng := len(prevTape)
    index := conf.Index

    // write the next symbol
    if (index == leng) && (next_symbol == Blank) {
        next_conf.Tape = prevTape
    } else if (index == leng) && (next_symbol != Blank) {
        next_conf.Tape = append(prevTape, next_symbol)
    } else if (index == 0) && (next_symbol == Blank) {
        // replace first symbol with a Blank and DO NOT move
        next_conf.Tape = append(prevTape[:0], prevTape[1:] ...)
        next_conf.Index = index
        return next_conf
    } else {
        next_conf.Tape = prevTape
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
