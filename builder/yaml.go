// Package yaml provides a function to translate a YAML file to a Turing machine.
package yaml

import (

    "gopkg.in/yaml.v2"
    "github.com/cjcodell1/tint/machine/turing"
    "github.com/cjcodell1/tint/file"
)

// tmBuilder is the struct to use to marshal the YAML.
type tmBuilder struct {
    // These must be exported, yaml parser requires it.
    Start string
    Accept string
    Reject string
    Transitions [][5]string
}

// Build creates a Turing machine from a YAML file.
func Build(configPath string) (turing.TuringMachine, error) {

    config, errRead := file.ReadAll(configPath)
    if errRead != nil {
        return nil, errRead
    }

    var builder tmBuilder

    errUnMarsh := yaml.Unmarshal([]byte(config), &builder)
    if errUnMarsh != nil {
        return nil, errUnMarsh
    }

    var trans []turing.Transition
    for _, t := range builder.Transitions {
        trans = append(trans, turing.Transition{turing.Input{t[0], t[1]}, turing.Output{t[2], t[3], t[4]}})
    }

    tm, errBuild := turing.NewTuringMachine(trans, builder.Start, builder.Accept, builder.Reject)
    if errBuild != nil {
        return nil, errBuild
    }

    return tm, nil

}
