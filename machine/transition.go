// Package for all machines.
package machine

// Interface for all types of Transitions.
type Transition interface {
	// Returns the inputs of the Transition.
	GetInput() []string
	// Returns the outputs of the Transition.
	GetOutput() []string
	// Checks if the given inputs are the inputs of the Transition.
	IsInput(inputs []string) (bool, error)
	// Checks if the given outputs are the outputs of the Transition.
	IsOutput(inputs []string) (bool, error)
}
