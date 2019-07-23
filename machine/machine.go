package machine


// interface for all Machines (e.g. DFA, NFA, PDA, various TMs, etc.)
type Machine interface {
	Start(input string)
	Step(conf Config) (Config, error)
	IsAccept(conf Config) bool
	IsReject(conf Config) bool
}
