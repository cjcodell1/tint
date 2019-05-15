package tm

const (
	Left string = "L"
	Right string = "R"
)

type Transition struct {
	In Input
	Out Output
}

type Input struct {
	State string
	Symbol string
}

type Output struct {
	State string
	Symbol string
	Move string
}
