package finite

import (
	"strings"
)

type Config struct {
	State string
	String []string
}

func (conf Config) Print() string {
	var line strings.Builder

	// the WriteString method on a strings.Builder always returns a nil error
	line.WriteString(conf.State)
	line.WriteString(" ")
	line.WriteString(strings.Join(conf.String, " "))
	return line.String()
}
