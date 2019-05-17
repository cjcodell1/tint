package yaml_builder

import (

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

    var builder tmBuilder

    errUnMarsh := yaml.Unmarshal([]byte(config), &builder)
    if errUnMarsh != nil {
        return nil, errUnMarsh
    }

    var trans []tm.Transition
    for _, t := range builder.Transitions {
        trans = append(trans, tm.Transition{tm.Input{t[0], t[1]}, tm.Output{t[2], t[3], t[4]}})
    }

    tm := tm.NewTuringMachine(trans, builder.Start, builder.Accept, builder.Reject)
    //tm, errBuild := tm.NewTuringMachine(trans, builder.Start, builder.Accept, builder.Reject)
    //if errBuild != nil {
    //    return nil, errBuild
    //}

    return tm, nil

}
