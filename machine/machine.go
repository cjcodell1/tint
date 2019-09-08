// Package for all machines.
package machine

// represents the available types of machines
const (
	DFA = "dfa"
	TM  = "tm"
)

const (
	Wildcard string = "*"
)

// interface for all Machines (e.g. DFA, NFA, PDA, various TMs, etc.)
type Machine interface {
	Start(input string) Configuration
	Step(conf Configuration) (Configuration, error)
	IsAccept(conf Configuration) bool
	IsReject(conf Configuration) bool
}
