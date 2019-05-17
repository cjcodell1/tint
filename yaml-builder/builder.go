package yaml_builder

import (

    "gopkg.in/yaml.v3"
    "github.com/cjcodell1/tint/tm"
    "github.com/cjcodell1/tint/file_reader"
)

type tmBuilder struct {
    start string
    accept string
    reject string
    transitions [][5]string
}

func Build(configPath string) tm.TuringMachine, error{

    config, err file_reader.ReadAll(configPath)
    if err != nil {
        return nil, err
    }

    builder := tmBuilder{}

    err := yaml.Unmarshal([]byte(config), &builder)
    if err != nil {
        return nil, err
    }

    tm, err := tm.NewTuringMachine(builder.transitions, builder.start, builder.accept, builder.reject)
    if err != nil {
        return nil, err
    }

    return tm, nil

}
