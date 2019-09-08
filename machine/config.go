// Package for all machines
package machine

// Interface for all types of configurations of machines.
type Configuration interface {
	// Prints a string representation of the Configuration.
	Print() string
	// Checks if the given string is the state of the Configuration.
	IsState(state string) bool
	// Determines if there can be a next Configuration.
	CanNext() bool
	// Computes the next Configuration given the next state, symbol, etc.
	// error if this function would produce an illegal Configuration.
	Next(inputs []string) (Configuration, error)
	// Gets the important information from a Configuration to find/perform the next Transition
	GetNext() ([]string, error)
}
