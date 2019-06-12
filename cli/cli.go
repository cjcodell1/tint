package cli

import (
    "fmt"
    "flag"
    "log"
    "strings"

    "github.com/cjcodell1/tint/tm"
    "github.com/cjcodell1/tint/file_reader"
    "github.com/cjcodell1/tint/yaml_builder"
)


var verboseFlag bool // prints out the step-by-step simulation
var testFlag bool // gives the machine a single test instead of a file of tests

/*
The test flag is a boolean because I think it is easier for the user to
provide the test after the Turing machine file, instead of before.
*/


func init() {
    const (
        usage = "print out the step-by-step simulation"
    )
    flag.BoolVar(&verboseFlag, "verbose", false, usage)
    flag.BoolVar(&verboseFlag, "v", false, usage + " (short-hand)")
}

func init() {
    const (
        usage = "provide a test to simulate on the the Turing machine (in place of a file of tests)"
    )
    flag.BoolVar(&testFlag, "test", false, usage)
    flag.BoolVar(&testFlag, "t", false, usage + " (short-hand)")
}

func Run() {
    if !flag.Parsed() {
        flag.Parse()
    }

    // check args
    if len(flag.Args()) != 2 {
        flag.PrintDefaults()
        log.Fatalln("Please input a Turing machine and test(s).")
    }

    var machine tm.TuringMachine
    var tests []string

    // get the Turing machine
    tmPath := flag.Arg(0)
    machine, err := yaml_builder.Build(tmPath)
    if err != nil {
        flag.PrintDefaults()
        log.Fatalln(err.Error())
    }
    // get the list of tests
    if testFlag {
        test := flag.Arg(1)
        tests = append(tests, test)
    } else {
        testsPath := flag.Arg(1)
        tests, err = file_reader.Lines(testsPath)
        if err != nil {
            flag.PrintDefaults()
            log.Fatalln(err.Error())
        }
    }

    // start TM with test
    for _, input := range tests {
        fmt.Printf("Simulating with %q.\n", input)
        var conf tm.Config
        for conf = machine.Start(input); !(machine.IsAccept(conf) || machine.IsReject(conf)); conf, err = machine.Step(conf) {
            if verboseFlag {
                fmt.Println(simplePrintConf(conf))
            }
        }
        if machine.IsAccept(conf) {
            fmt.Println("Accepted.\n")
        } else {
            fmt.Println("Rejected.\n")
        }
    }
}

func simplePrintConf(conf tm.Config) string {
    return fmt.Sprintf("%s: %q, at %d", conf.State, strings.Join(conf.Tape, " "), conf.Index)
}
