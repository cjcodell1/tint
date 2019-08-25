package machine

const (
	Wildcard string = "*"
)

// interface for all Machines (e.g. DFA, NFA, PDA, various TMs, etc.)
type Machine interface {
	Start(input string) Config
	Step(conf Config) (Config, error)
	IsAccept(conf Config) (bool, error)
	IsReject(conf Config) (bool, error)
}
