package two_test

import (
	testing

	github.com/cjcodell1/tint/machine
	github.com/cjcodell1/tint/machine/turing
	github.com/cjcodell1/tint/machine/turing/ways
	github.com/cjcodell1/tint/machine/turing/ways/two
)

// MakeTuringMachine
type makeTuringMachine struct {
	trans []ways.Transition
	start string
	accept string
	reject string
}

// Start
type start struct {
	tm Turing.TuringMachine
	tmName string
	input string
	expect string
}

// Step
type step struct {
	tm Turing.TuringMachine
	tmName string
	input two.Config
	expect string
}

// IsAccept
type isAccept struct [
	tm Turing.Machine
	tmName string
	input two.Config
}

// IsReject
type isReject struct [
	tm Turing.Machine
	tmName string
	input two.Config
}

var (
	makeTuringMachineTests makeTuringMachine
	startTests start
	stepTests step
	isAcceptTests isAccept
	isRejectTests isReject
)

var (
	emptyTM machine.Machine
	allTM machine.Machine
	addMarkersTM machine.Machine
	moveTM machine.Machine
	addBlankTM machine.Machine
	starSymTM machine.Machine
	brokenSymTM machine.Machine
	starStateTM machine.Machine
	brokenStateTM machine.Machine
	doNothingTM machine.Machine
)

func TestNewTuringMachine(t *testing.T) {
	for _, tc := range makeTuringMachineTests {
		got, _ := two.MakeTuringMachine(tc.trans, tc.start, tc.accept, tc.reject)
		_, ok := got.(machine.Machine)
		if !ok {
			t.Error("Did not create a Turing Machine.")
		}
	}
}

func TestStart(t *testing.T) {
	for _, tc := range startTests {
		got := tc.tm.Start(tc.input)
		if got != tc.expected {
			t.Errorf("%s.Start(%s) == %s != %s", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestStep(t *testing.T) {
	for _, tc := range stepTests {
		got, _ := tc.tm.Step(tc.input)
		if got != tc.expected {
			t.Errorf("%s.Step(%q) == %q != %q", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestIsAccept(t *testing.T) {
	for _, tc := range isAcceptTests {
		got := tc.tm.IsAccept(tc.input)
		if got != tc.expected {
			t.Errorf("%s.IsAccept(%q) == %s != %s", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestIsReject(t *testing.t) {
	for _, tc := range isRejectTests {
		got := tc.tm.IsReject(tc.input)
		if got != tc.expected {
			t.Errorf("%s.IsReject(%q) == %s != %s", tc.tmName, tc.input, got, tc.expect)
		}
	}
}


// Turing machines to test
emptyTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("start", "a", "reject", "c", ways.Right),
		ways.MakeTransition("start", "b", "reject", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

allTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("start", "a", "accept", "c", ways.Right),
		ways.MakeTransition("start", "b", "accept", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

addMarkersTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("q0", "a", "placeA", "$", ways.Right),
		ways.MakeTransition("q0", "b", "placeB", "$", ways.Right),
		ways.MakeTransition("q0", turing.Blank, "place$", "$", ways.Right),

		ways.MakeTransition("placeA", "a", "placeA", "a", ways.Right),
		ways.MakeTransition("placeA", "b", "placeB", "a", ways.Right),
		ways.MakeTransition("placeA", turing.Blank, "place$", "a", ways.Right),

		ways.MakeTransition("placeB", "a", "placeA", "b", ways.Right),
		ways.MakeTransition("placeB", "b", "placeB", "b", ways.Right),
		ways.MakeTransition("placeB", turing.Blank, "place$", "b", ways.Right),

		ways.MakeTransition("place$", turing.Blank, "done", "$", ways.Right),
	},
	"q0",
	"done",
	"cannot_reject")

moveTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("q0", "l", "q0", "l", ways.Left),
		ways.MakeTransition("q0", "r", "q0", "r", ways.Right),
	},
	"q0",
	"accept",
	"reject")

placeBlankTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("q0", "l", "q0", turing.Blank, ways.Left),
		ways.MakeTransition("q0", "r", "q0", turing.Blank, ways.Right),
	},
	"q0",
	"accept",
	"reject")

starSymTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("next", turing.Blank, "accept", turing.Blank, ways.Right),
		ways.MakeTransition("next", "*", "next", "*", ways.Right),
	},
	"next",
	"accept",
	"reject")

brokenSymTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("next", "*", "next", "*", ways.Right),
		ways.MakeTransition("next", turing.Blank, "accept", turing.Blank, ways.Right),
	},
	"next",
	"accept",
	"reject")

starStateTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("*", turing.Blank, "accept", turing.Blank, ways.Right),
		ways.MakeTransition("*", "c", "*", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

brokenStateTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("*", "c", "*", "c", ways.Right),
		ways.MakeTransition("*", turing.Blank, "accept", turing.Blank, ways.Right),
	},
	"start",
	"accept",
	"reject")

doNothingTM = MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("*", "*", "*", "*", ways.Right),
	},
	"start",
	"accept",
	"reject")

// set up the makeTuringMachineTests automatically
func init() {
	makeTuringMachineTests = []makeTuringMachine{
		{[]ways.Transition{}, "start", "accept", "reject"},
		{[]ways.Transition{ways.MakeTransition{"q0", "*", "accept", "*", turing.Right}},
	}
}

// set up the startTests automatically
func init() {
	startTests = []start{
		{emptyTM, "emptyTM", "a b a", "{start {a b a} 0}"},
		{allTM, "allTM", "", "{start {} 0}"},
		{addMarkersTM, "addMarkersTM", "a a a a a", "{q0 {a a a a a} 0}"},
	}
}

// set up the stepTests automatically
func init() {
	var (
		start two.Config
		step1 two.Config
		step2 two.Config
		step3 two.Config
		step4 two.Config
		step5 two.Config
		step6 two.Config
		step7 two.Config
		step8 two.Config
	)

	// adding the tests for emptyTM
	start = emptyTM.Start("a b")
	step1, _ = emptyTM.Step(start)
	step2, _ = emptyTM.Step(step1)
	step3, _ = emptyTM.Step(step2)
	stepTests = append(stepTests, []step{
		{emptyTM, "emptyTM", step1, "{start, {a b}, 0}"},
		{emptyTM, "emptyTM", step2, "{reject, {c b}, 1}"},
		{emptyTM, "emptyTM", step3, "{reject, {c b}, 1}"},
	})

	// adding the tests for allTM
	start = allTM.Start("b a")
	step1, _ = allTM.Step(start)
	step2, _ = allTM.Step(step1)
	step3, _ = allTM.Step(step2)
	stepTests = append(stepTests, []step{
		{allTM, "allTM", step1, "{start, {b a}, 0}"},
		{allTM, "allTM", step2, "{accept, {c a}, 1}"},
		{allTM, "allTM", step3, "{accept, {c a}, 1}"},
	})

	// repeat for the rest of the Turing machines -- just set the steps and then append the new tests to the stepTests
}
