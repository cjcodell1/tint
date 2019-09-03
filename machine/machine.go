// Package for all machines.
package machine

// represents the available types of machines
const (
	DFA = "dfa"
	TM = "tm"
)

const (
	Wildcard string = "*"
)

// interface for all Machines (e.g. DFA, NFA, PDA, various TMs, etc.)
type Machine interface {
	Start(input string) Config
	Step(conf Config) (Config, error)
	IsAccept(conf Config) bool
	IsReject(conf Config) bool
}
