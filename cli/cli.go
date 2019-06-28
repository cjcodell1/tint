// Package cli provides a small command line interface for this program.
package cli

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/cjcodell1/tint/file"
	"github.com/cjcodell1/tint/tm"
	"github.com/cjcodell1/tint/worker"
	"github.com/cjcodell1/tint/yaml_builder"
)

var (
	verboseFlag bool // prints out the step-by-step simulation
	testFlag    bool // use a single test instead of a file of tests
)

func init() {
	const (
		usage = "print out the step-by-step simulation"
	)
	flag.BoolVar(&verboseFlag, "verbose", false, usage)
	flag.BoolVar(&verboseFlag, "v", false, usage+" (short-hand)")
}

func init() {
	const (
		usage = "provide a test to simulate on the the Turing machine (in place of a file of tests)"
	)
	flag.BoolVar(&testFlag, "test", false, usage)
	flag.BoolVar(&testFlag, "t", false, usage+" (short-hand)")
}

// Run starts the program by building the Turing machine and
// simulating it with test(s).
func Run() {
	// Ensures the flags are parsed.
	if !flag.Parsed() {
		flag.Parse()
	}

	// Ensures there are two non-flag arguments.
	if len(flag.Args()) != 2 {
		flag.PrintDefaults()
		log.Fatalln("Please input a Turing machine and test(s).")
	}

	var machine tm.TuringMachine
	var tests []string

	// Builds the Turing machine from the first non-flag argument.
	tmPath := flag.Arg(0)
	machine, err := yaml_builder.Build(tmPath)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalln(err.Error())
	}

	// Builds the slice of tests used for testing from the second non-flag argument.
	if testFlag {
		test := flag.Arg(1)
		tests = append(tests, test)
	} else {
		testsPath := flag.Arg(1)
		tests, err = file.ReadLines(testsPath)
		if err != nil {
			flag.PrintDefaults()
			log.Fatalln(err.Error())
		}
	}

	// Simulates the tests and prints the results.
	results := worker.TestAll(tests, machine)
	for _, result := range results {
		fmt.Printf("Simulating with %q.\n", result.Input)
		if verboseFlag {
			for _, conf := range result.Configs {
				fmt.Println(simplePrintConf(conf))
			}
		}
		if result.TM.IsAccept(result.Configs[len(result.Configs)-1]) {
			fmt.Println("Accepted.\n")
		} else {
			fmt.Println("Rejected.\n")
		}
	}
}

func simplePrintConf(conf tm.Config) string {
	return fmt.Sprintf("%s: %q, at %d", conf.State, strings.Join(conf.Tape, " "), conf.Index)
}
