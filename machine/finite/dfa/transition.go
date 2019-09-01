package dfa

type Transition struct {
	in input
	out output
}

type input struct {
	state string
	symbol string
}

type output struct {
	symbol string
}

func MakeTransition(in_state string, in_symbol string, out_state string) {
	return Transition{input{in_state, in_symbol}, output{out_state}}
}

func (t Transition) GetInput() (string, string) {
	return t.in.state, t.in.symbol
}

func (t Transition) GetOutput() string {
	return t.out.state
}
