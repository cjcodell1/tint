package yaml_builder

import (
	"fmt"
	"strings"

    "gopkg.in/yaml.v2"
    "github.com/cjcodell1/tint/tm"
    "github.com/cjcodell1/tint/file_reader"
)

type tmBuilder struct {
    Start string
    Accept string
    Reject string
    Transitions [][5]string
}

func Build(configPath string) (tm.TuringMachine, error) {

    config, errRead := file_reader.ReadAll(configPath)
    if errRead != nil {
        return nil, errRead
    }

	fmt.Println(config)
    var builder tmBuilder

    errUnMarsh := yaml.Unmarshal([]byte(config), &builder)
    if errUnMarsh != nil {
        return nil, errUnMarsh
    }

	fmt.Println(builder.Start)
	fmt.Println(builder.Accept)
	fmt.Println(builder.Reject)
	for _, t := range builder.Transitions {
		fmt.Println(strings.Join(t[:], " "))
	}


    var trans []tm.Transition
    for i, t := range builder.Transitions {
        trans[i] = tm.Transition{tm.Input{t[0], builder.Transitions[i][1]}, tm.Output{t[2], builder.Transitions[i][3], builder.Transitions[i][4]}}
    }

    tm := tm.NewTuringMachine(trans, builder.Start, builder.Accept, builder.Reject)
    //tm, errBuild := tm.NewTuringMachine(trans, builder.Start, builder.Accept, builder.Reject)
    //if errBuild != nil {
    //    return nil, errBuild
    //}

    return tm, nil

}
