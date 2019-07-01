// Packing turing is the implementation of a Turing machine.
package turing

// Config is a Turing machine configuration.
type Config struct {
	State string
	Tape  []string
	Index int
}
