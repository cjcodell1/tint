package yaml-builder

import (
	"log"

	"gopkg.in/yaml.v3"
	"github.com/cjcodell1/tint/tm"
)

type TMBuilder struct {
	Start string
	Accept string
	Reject string
	Transitions [][5]string
}

func build(config string) tm.TuringMachine {
	builder := TMBuilder{}

	err := yaml.Unmarshal([]byte(config), &builder)
	if err != nil {
		log.Fatal(err)
	}

	// OK -- now I've got a Builder I need to use
	// I want to do some ensurance tests (checking the accept state != reject)
	// So I need to create a NewTuringMachine function in my turing machine file.
	// I can not export the struct (lower case), but still export the fields (keeping the caps)
	// Use this function to return the actual TuringMachine 
}
