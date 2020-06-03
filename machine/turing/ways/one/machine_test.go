package one_test

import (
	"fmt"
	"testing"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/turing"
	"github.com/cjcodell1/tint/machine/turing/ways/one"
)

// MakeTuringMachine
type makeTuringMachine struct {
	trans  [][]string
	start  string
	accept string
	reject string
}

// Start
type start struct {
	tm     machine.Machine
	tmName string
	input  string
	expect string
}

// Step
type step struct {
	tm     machine.Machine
	tmName string
	input  machine.Configuration
	expect string
}

// IsAccept
type isAccept struct {
	tm     machine.Machine
	tmName string
	input  machine.Configuration
	expect bool
}

// IsReject
type isReject struct {
	tm     machine.Machine
	tmName string
	input  machine.Configuration
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
		got, _ := one.MakeTuringMachine(tc.trans, tc.start, tc.accept, tc.reject)
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
		got := tc.tm.IsAccept(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsAccept(%v) == %t != %t", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

func TestIsReject(t *testing.T) {
	for _, tc := range isRejectTests {
		got := tc.tm.IsReject(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsReject(%v) == %t != %t", tc.tmName, tc.input, got, tc.expect)
		}
	}
}

// Turing machines to test
var emptyTM, _ = one.MakeTuringMachine(
	[][]string{
		{"start", "a", "reject", "c", turing.Right},
		{"start", "b", "reject", "c", turing.Right},
	},
	"start",
	"accept",
	"reject")

var allTM, _ = one.MakeTuringMachine(
	[][]string{
		{"start", "a", "accept", "c", turing.Right},
		{"start", "b", "accept", "c", turing.Right},
	},
	"start",
	"accept",
	"reject")

var addMarkersTM, _ = one.MakeTuringMachine(
	[][]string{
		{"q0", "a", "placeA", "$", turing.Right},
		{"q0", "b", "placeB", "$", turing.Right},
		{"q0", turing.Blank, "place$", "$", turing.Right},

		{"placeA", "a", "placeA", "a", turing.Right},
		{"placeA", "b", "placeB", "a", turing.Right},
		{"placeA", turing.Blank, "place$", "a", turing.Right},

		{"placeB", "a", "placeA", "b", turing.Right},
		{"placeB", "b", "placeB", "b", turing.Right},
		{"placeB", turing.Blank, "place$", "b", turing.Right},

		{"place$", turing.Blank, "done", "$", turing.Right},
	},
	"q0",
	"done",
	"cannot_reject")

var moveTM, _ = one.MakeTuringMachine(
	[][]string{
		{"q0", "l", "q0", "l", turing.Left},
		{"q0", "r", "q0", "r", turing.Right},
		{"q0", turing.Blank, "q0", "r", turing.Right},
	},
	"q0",
	"accept",
	"reject")

var placeBlankTM, _ = one.MakeTuringMachine(
	[][]string{
		{"q0", "l", "q0", turing.Blank, turing.Left},
		{"q0", "r", "q0", turing.Blank, turing.Right},
		{"q0", turing.Blank, "q0", turing.Blank, turing.Right},
	},
	"q0",
	"accept",
	"reject")

var starSymTM, _ = one.MakeTuringMachine(
	[][]string{
		{"next", turing.Blank, "accept", turing.Blank, turing.Right},
		{"next", "*", "next", "*", turing.Right},
	},
	"next",
	"accept",
	"reject")

var brokenSymTM, _ = one.MakeTuringMachine(
	[][]string{
		{"next", "*", "next", "*", turing.Right},
		{"next", turing.Blank, "accept", turing.Blank, turing.Right},
	},
	"next",
	"accept",
	"reject")

var starStateTM, _ = one.MakeTuringMachine(
	[][]string{
		{"*", turing.Blank, "accept", turing.Blank, turing.Right},
		{"*", "a", "*", "c", turing.Right},
	},
	"start",
	"accept",
	"reject")

var doNothingTM, _ = one.MakeTuringMachine(
	[][]string{
		{"*", "*", "*", "*", turing.Right},
	},
	"start",
	"accept",
	"reject")

// set up the makeTuringMachineTests automatically
func init() {
	makeTuringMachineTests = []makeTuringMachine{
		{[][]string{}, "start", "accept", "reject"},
		{[][]string{{"q0", "*", "accept", "*", turing.Right}}, "start", "accept", "reject"},
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
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration
	var step3 machine.Configuration
	var step4 machine.Configuration

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
	step1, _ = moveTM.Step(start)
	step2, _ = moveTM.Step(step1)
	stepTests = append(stepTests, []step{
		{moveTM, "moveTM", start, "{q0 [r] 1}"},
		{moveTM, "moveTM", step1, "{q0 [r r] 2}"},
		{moveTM, "moveTM", step2, "{q0 [r r r] 3}"},
	}...)

	// adding the tests for placeBlankTM
	start = placeBlankTM.Start("l")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [_] 0}"},
	}...)

	start = placeBlankTM.Start("r")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [_] 1}"},
	}...)

	start = placeBlankTM.Start("")
	stepTests = append(stepTests, []step{
		{placeBlankTM, "placeBlankTM", start, "{q0 [_] 1}"},
	}...)

	// adding the tests for starSymTM
	start = starSymTM.Start("a a")
	step1, _ = starSymTM.Step(start)
	step2, _ = starSymTM.Step(step1)
	stepTests = append(stepTests, []step{
		{starSymTM, "starSymTM", start, "{next [a a] 1}"},
		{starSymTM, "starSymTM", step1, "{next [a a] 2}"},
		{starSymTM, "starSymTM", step2, "{accept [a a _] 3}"},
	}...)

	// adding the tests for brokenSymTM
	start = brokenSymTM.Start("a a")
	step1, _ = brokenSymTM.Step(start)
	step2, _ = brokenSymTM.Step(step1)
	stepTests = append(stepTests, []step{
		{brokenSymTM, "brokenSymTM", start, "{next [a a] 1}"},
		{brokenSymTM, "brokenSymTM", step1, "{next [a a] 2}"},
		{brokenSymTM, "brokenSymTM", step2, "{next [a a _] 3}"},
	}...)

	// adding the tests for starStateTM
	start = starStateTM.Start("a a")
	step1, _ = starStateTM.Step(start)
	step2, _ = starStateTM.Step(step1)
	stepTests = append(stepTests, []step{
		{starStateTM, "starStateTM", start, "{start [c a] 1}"},
		{starStateTM, "starStateTM", step1, "{start [c c] 2}"},
		{starStateTM, "starStateTM", step2, "{accept [c c _] 3}"},
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
		{doNothingTM, "doNothingTM", step3, "{start [hello world !!! _] 4}"},
	}...)
}

// set up the isAcceptTests automatically
func init() {
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration
	var step3 machine.Configuration

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
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration

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
