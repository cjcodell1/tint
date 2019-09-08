// Package cli provides a small command line interface for this program.
package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cjcodell1/tint/builder/yaml"
	"github.com/cjcodell1/tint/file"
	"github.com/cjcodell1/tint/machine"
)

var (
	verboseFlag bool   // prints out the step-by-step simulation
	testFlag    bool   // use a single test instead of a file of tests
	machineFlag string // denotes what type of machine is specified
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
		usage = "provide a test to simulate on the machine (in place of a file of tests)"
	)
	flag.BoolVar(&testFlag, "test", false, usage)
	flag.BoolVar(&testFlag, "t", false, usage+" (short-hand)")
}

func init() {
	const (
		usage = "denote what type of machine is specified"
	)
	flag.StringVar(&machineFlag, "machine", "", usage)
	flag.StringVar(&machineFlag, "m", "", usage+" (short-hand)")
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
		fmt.Println("Please provide the machine and test(s).")
		os.Exit(1)
	}

	// Ensures the machine flag was set.
	if machineFlag == "" {
		flag.PrintDefaults()
		fmt.Println("Please provide the type of machine the file specifies.")
		os.Exit(1)
	}
	// normalizes the machine flag
	machineFlag = strings.ToLower(machineFlag)
	var m machine.Machine
	var tests []string

	// Builds the Turing machine from the first non-flag argument.
	mPath := flag.Arg(0)
	m, err := yaml.Build(mPath, machineFlag)
	if err != nil {
		flag.PrintDefaults()
		fmt.Println("There was an error building your machine.")
		fmt.Println(err)
		os.Exit(1)
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
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Simulate the test
	totalAccept := 0
	totalReject := 0
	totalError := 0
	for _, input := range tests {
		fmt.Printf("Simulating with \"%s\".\n", input)

		conf := m.Start(input)
		for {
			// print verbosely
			if verboseFlag {
				fmt.Println(conf.Print())
			}

			// check if accept or reject and break
			if m.IsAccept(conf) {
				totalAccept += 1
				fmt.Println("Accepted.\n")
				break
			} else if m.IsReject(conf) {
				totalReject += 1
				fmt.Println("Rejected.\n")
				break
			}

			// step
			conf, err = m.Step(conf)
			if err != nil {
				fmt.Println("ERROR! Please see below:")
				fmt.Println(err)
				fmt.Println("Skipping this test.\n")
				totalError += 1
				break
			}
		}
	}
	fmt.Printf("%d accepted.\n", totalAccept)
	fmt.Printf("%d rejected.\n", totalReject)
	fmt.Printf("%d errors.\n", totalError)
}
