package ways

const (
	Left  string = "L"
	Right string = "R"
)

// Transition represents a transition function.
type Transition struct {
	in  input
	out output
}

// Input represents an input to a transition function.
type input struct {
	state  string
	symbol string
}

// Output represents an input to a transiiton function.
type output struct {
	state  string
	symbol string
	move   string
}

func MakeTransition(in_state string, in_sym string, out_state string, out_sym string, out_move string) Transition {
	return Transition{input{in_state, in_sym}, output{out_state, out_sym, out_move}}
}

func (t Transition) GetInput() (string, string) {
	return t.in.state, t.in.symbol
}

func (t Transition) GetOutput() (string, string, string) {
	return t.out.state, t.out.symbol, t.out.move
}
