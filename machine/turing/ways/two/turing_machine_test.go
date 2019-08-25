package two_test

import (
	"testing"
	"fmt"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
	"github.com/cjcodell1/tint/machine/turing/ways"
	"github.com/cjcodell1/tint/machine/turing/ways/two"
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
	tm machine.Machine
	tmName string
	input string
	expect string
}

// Step
type step struct {
	tm machine.Machine
	tmName string
	input machine.Config
	expect string
}

// IsAccept
type isAccept struct {
	tm machine.Machine
	tmName string
	input machine.Config
	expect bool
}

// IsReject
type isReject struct {
	tm machine.Machine
	tmName string
	input machine.Config
	expect bool
}

var makeTuringMachineTests []makeTuringMachine
var startTests []start
var stepTests []step
var isAcceptTests []isAccept
var isRejectTests []isReject

// var emptyTM machine.Machine
// var allTM machine.Machine
// var addMarkersTM machine.Machine
// var moveTM machine.Machine
// var addBlankTM machine.Machine
// var starSymTM machine.Machine
// var brokenSymTM machine.Machine
// var starStateTM machine.Machine
// var doNothingTM machine.Machine

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
		got := fmt.Sprint(tc.tm.Start(tc.input))
		if got != tc.expect {
			t.Errorf("%s.Start(%s) == %v != %s", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestStep(t *testing.T) {
	for _, tc := range stepTests {
		gotMachine, _ := tc.tm.Step(tc.input)
		got := fmt.Sprint(gotMachine)
		if got != tc.expect {
			t.Errorf("%s.Step(%v) == %v != %s", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestIsAccept(t *testing.T) {
	for _, tc := range isAcceptTests {
		got, _ := tc.tm.IsAccept(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsAccept(%v) == %t != %t", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestIsReject(t *testing.T) {
	for _, tc := range isRejectTests {
		got, _ := tc.tm.IsReject(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsReject(%v) == %t != %t", tc.tmName, tc.input, got, tc.expect)
		}
	}
}


// Turing machines to test
var emptyTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("start", "a", "reject", "c", ways.Right),
		ways.MakeTransition("start", "b", "reject", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

var allTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("start", "a", "accept", "c", ways.Right),
		ways.MakeTransition("start", "b", "accept", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

var addMarkersTM, _ = two.MakeTuringMachine(
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

var moveTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("q0", "l", "q0", "l", ways.Left),
		ways.MakeTransition("q0", "r", "q0", "r", ways.Right),
		ways.MakeTransition("q0", turing.Blank, "q0", "r", ways.Right),
	},
	"q0",
	"accept",
	"reject")

var placeBlankTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("q0", "l", "q0", turing.Blank, ways.Left),
		ways.MakeTransition("q0", "r", "q0", turing.Blank, ways.Right),
		ways.MakeTransition("q0", turing.Blank, "q0", turing.Blank, ways.Right),
	},
	"q0",
	"accept",
	"reject")

var starSymTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("next", turing.Blank, "accept", turing.Blank, ways.Right),
		ways.MakeTransition("next", "*", "next", "*", ways.Right),
	},
	"next",
	"accept",
	"reject")

var brokenSymTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("next", "*", "next", "*", ways.Right),
		ways.MakeTransition("next", turing.Blank, "accept", turing.Blank, ways.Right),
	},
	"next",
	"accept",
	"reject")

var starStateTM, _ = two.MakeTuringMachine(
	[]ways.Transition{
		ways.MakeTransition("*", turing.Blank, "accept", turing.Blank, ways.Right),
		ways.MakeTransition("*", "a", "*", "c", ways.Right),
	},
	"start",
	"accept",
	"reject")

var doNothingTM, _ = two.MakeTuringMachine(
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
		{[]ways.Transition{ways.MakeTransition("q0", "*", "accept", "*", ways.Right)}, "start", "accept", "reject"},
	}
}

// set up the startTests automatically
func init() {
	startTests = []start{
		{emptyTM, "emptyTM", "a b a", "{start [a b a] 0}"},
		{allTM, "allTM", "", "{start [] 0}"},
		{addMarkersTM, "addMarkersTM", "a a a a a", "{q0 [a a a a a] 0}"},
	}
}

// set up the stepTests automatically
func init() {
	var start machine.Config
	var step1 machine.Config
	var step2 machine.Config
	var step3 machine.Config
	var step4 machine.Config

	// adding the tests for emptyTM
	start = emptyTM.Start("a b")
	step1, _ = emptyTM.Step(start)
	stepTests = append(stepTests, []step{
		{emptyTM, "emptyTM", start, "{reject [c b] 1}"},
		{emptyTM, "emptyTM", step1, "{reject [c b] 1}"},
	}...)

	// adding the tests for allTM
	start = allTM.Start("b a")
	step1, _ = allTM.Step(start)
	stepTests = append(stepTests, []step{
		{allTM, "allTM", start, "{accept [c a] 1}"},
		{allTM, "allTM", step1, "{accept [c a] 1}"},
	}...)

	// adding the tests for addMarkersTM
	start = addMarkersTM.Start("")
	step1, _ = addMarkersTM.Step(start)
	stepTests = append(stepTests, []step{
		{addMarkersTM, "addMarkersTM", start, "{place$ [$] 1}"},
		{addMarkersTM, "addMarkersTM", step1, "{done [$ $] 2}"},
	}...)

	start = addMarkersTM.Start("b a b")
	step1, _ = addMarkersTM.Step(start)
	step2, _ = addMarkersTM.Step(step1)
	step3, _ = addMarkersTM.Step(step2)
	step4, _ = addMarkersTM.Step(step3)
	stepTests = append(stepTests, []step{
		{addMarkersTM, "addMarkersTM", start, "{placeB [$ a b] 1}"},
		{addMarkersTM, "addMarkersTM", step1, "{placeA [$ b b] 2}"},
		{addMarkersTM, "addMarkersTM", step2, "{placeB [$ b a] 3}"},
		{addMarkersTM, "addMarkersTM", step3, "{place$ [$ b a b] 4}"},
		{addMarkersTM, "addMarkersTM", step4, "{done [$ b a b $] 5}"},
	}...)

	// adding the tests for moveTM
	start = moveTM.Start("l")
	step1, _ = moveTM.Step(start)
	step2, _ = moveTM.Step(step1)
	stepTests = append(stepTests, []step{
		{moveTM, "moveTM", start, "{q0 [l] 0}"},
		{moveTM, "moveTM", step1, "{q0 [l] 0}"},
		{moveTM, "moveTM", step2, "{q0 [l] 0}"},
	}...)

	start = moveTM.Start("r")
	//fmt.Printf("start: %v\n--------\n", start)
	step1, _ = moveTM.Step(start)
	//fmt.Printf("step1: %v\n--------\n", step1)
	step2, _ = moveTM.Step(step1)
	//fmt.Printf("step2: %v\n--------\n", step2)
	stepTests = append(stepTests, []step{
		{moveTM, "moveTM", start, "{q0 [r] 1}"},
		{moveTM, "moveTM", step1, "{q0 [r r] 2}"},
		{moveTM, "moveTM", step2, "{q0 [r r r] 3}"},
	}...)

	// adding the tests for placeBlankTM
	start = placeBlankTM.Start("l")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [] 0}"},
	}...)

	start = placeBlankTM.Start("r")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [] 0}"},
	}...)

	start = placeBlankTM.Start("")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [] 0}"},
	}...)

	// adding the tests for starSymTM
	start = starSymTM.Start("a a")
	step1, _ = starSymTM.Step(start)
	step2, _ = starSymTM.Step(step1)
	stepTests = append(stepTests, []step{
		{starSymTM, "starSymTM", start, "{next [a a] 1}"},
		{starSymTM, "starSymTM", step1, "{next [a a] 2}"},
		{starSymTM, "starSymTM", step2, "{accept [a a] 2}"},
	}...)

	// adding the tests for brokenSymTM
	start = brokenSymTM.Start("a a")
	step1, _ = brokenSymTM.Step(start)
	step2, _ = brokenSymTM.Step(step1)
	stepTests = append(stepTests, []step{
		{brokenSymTM, "brokenSymTM", start, "{next [a a] 1}"},
		{brokenSymTM, "brokenSymTM", step1, "{next [a a] 2}"},
		{brokenSymTM, "brokenSymTM", step2, "{next [a a] 2}"},
	}...)

	// adding the tests for starStateTM
	start = starStateTM.Start("a a")
	step1, _ = starStateTM.Step(start)
	step2, _ = starStateTM.Step(step1)
	stepTests = append(stepTests, []step{
		{starStateTM, "starStateTM", start, "{start [c a] 1}"},
		{starStateTM, "starStateTM", step1, "{start [c c] 2}"},
		{starStateTM, "starStateTM", step2, "{accept [c c] 2}"},
	}...)

	// adding the tests for doNothingTM
	start = doNothingTM.Start("hello world !!!")
	step1, _ = doNothingTM.Step(start)
	step2, _ = doNothingTM.Step(step1)
	step3, _ = doNothingTM.Step(step2)
	stepTests = append(stepTests, []step{
		{doNothingTM, "doNothingTM", start, "{start [hello world !!!] 1}"},
		{doNothingTM, "doNothingTM", step1, "{start [hello world !!!] 2}"},
		{doNothingTM, "doNothingTM", step2, "{start [hello world !!!] 3}"},
		{doNothingTM, "doNothingTM", step3, "{start [hello world !!!] 3}"},
	}...)
}

// set up the isAcceptTests automatically
func init() {
	var start machine.Config
	var step1 machine.Config
	var step2 machine.Config
	var step3 machine.Config

	// adding the tests for emptyTM
	start = emptyTM.Start("a")
	step1, _ = emptyTM.Step(start)
	isAcceptTests = append(isAcceptTests, []isAccept{
		{emptyTM, "emptyTM", start, false},
		{emptyTM, "emptyTM", step1, false},
	}...)

	// adding the tests for allTM
	start = allTM.Start("a")
	step1, _ = allTM.Step(start)
	isAcceptTests = append(isAcceptTests, []isAccept{
		{allTM, "allTM", start, false},
		{allTM, "allTM", step1, true},
	}...)

	// adding the tests for addMarkersTM
	start = addMarkersTM.Start("a")
	step1, _ = addMarkersTM.Step(start)
	step2, _ = addMarkersTM.Step(step1)
	step3, _ = addMarkersTM.Step(step2)
	isAcceptTests = append(isAcceptTests, []isAccept{
		{addMarkersTM, "addMarkersTM", start, false},
		{addMarkersTM, "addMarkersTM", step1, false},
		{addMarkersTM, "addMarkersTM", step2, false},
		{addMarkersTM, "addMarkersTM", step3, true},
	}...)
}

// set up the isRejectTests automatically
func init() {
	var start machine.Config
	var step1 machine.Config
	var step2 machine.Config

	// adding the tests for emptyTM
	start = emptyTM.Start("a")
	step1, _ = emptyTM.Step(start)
	isRejectTests = append(isRejectTests, []isReject{
		{emptyTM, "emptyTM", start, false},
		{emptyTM, "emptyTM", step1, true},
	}...)

	// adding the tests for allTM
	start = allTM.Start("a")
	step1, _ = allTM.Step(start)
	isRejectTests = append(isRejectTests, []isReject{
		{allTM, "allTM", start, false},
		{allTM, "allTM", step1, false},
	}...)

	// adding the tests for addMarkersTM
	start = addMarkersTM.Start("a")
	step1, _ = addMarkersTM.Step(start)
	step2, _ = addMarkersTM.Step(step1)
	isRejectTests = append(isRejectTests, []isReject{
		{addMarkersTM, "addMarkersTM", start, false},
		{addMarkersTM, "addMarkersTM", step1, false},
		{addMarkersTM, "addMarkersTM", step2, false},
	}...)
}
