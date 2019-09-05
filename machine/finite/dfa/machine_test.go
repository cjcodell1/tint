package dfa_test

import (
	"testing"
	"fmt"

	"github.com/cjcodell1/tint/machine"
	"github.com/cjcodell1/tint/machine/finite/dfa"
)

type makeDFAT struct {
	trans [][]string
	start string
	accepts []string
	err error
}

type startT struct {
	d machine.Machine
	name string
	input string
	expect string
}

type stepT struct {
	d machine.Machine
	name string
	input machine.Configuration
	expect string
	err error
}

type isAcceptT struct {
	d machine.Machine
	name string
	input machine.Configuration
	expect bool
}

type isRejectT struct {
	d machine.Machine
	name string
	input machine.Configuration
	expect bool
}

var makeDFATests []makeDFAT
var startTests []startT
var stepTests []stepT
var isAcceptTests []isAcceptT
var isRejectTests []isRejectT

func TestMakeDFA(t *testing.T) {
	for _, tc := range makeDFATests {
		got, err := dfa.MakeDFA(tc.trans, tc.start, tc.accepts)
		_, ok := got.(machine.Machine)
		if !ok {
			t.Error("Did not create a DFA.")
		}
		if err != tc.err {
			t.Errorf("Actual error %s != expected error %s", err, tc.err)
		}
	}
}

func TestStart(t *testing.T) {
	for _, tc := range startTests {
		got := fmt.Sprint(tc.d.Start(tc.input))
		if got != tc.expect {
			t.Errorf("%s.Start(%s) == %s != %s", tc.name, tc.input, got, tc.expect)
		}
	}
}

func TestStep(t *testing.T) {
	for _, tc := range stepTests {
		ans, err := tc.d.Step(tc.input)
		got := fmt.Sprint(ans)
		if got != tc.expect {
			t.Errorf("%s.Step(%s) == %s != %s", tc.name, tc.input, got, tc.expect)
		}
		if err != tc.err {
			t.Errorf("%s.Start(%s) == %s != %s", tc.name, tc.input, got, tc.expect)
		}
	}
}

func TestIsAccept(t *testing.T) {
	for _, tc := range isAcceptTests {
		got := tc.d.IsAccept(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsAccept(%s) == %t != %t", tc.name, tc.input, got, tc.expect)
		}
	}
}

func TestIsReject(t *testing.T) {
	for _, tc := range isRejectTests {
		got := tc.d.IsReject(tc.input)
		if got != tc.expect {
			t.Errorf("%s.IsReject(%s) == %t != %t", tc.name, tc.input, got, tc.expect)
		}
	}
}

var emptyDFA, _ = dfa.MakeDFA(
	[][]string{
		{"start", "a", "reject"},
		{"start", "b", "reject"},
		{"start", "c", "reject"},

		{"reject", "a", "reject"},
		{"reject", "b", "reject"},
		{"reject", "c", "reject"},
	},
	"start",
	[]string{})

var allDFA, _ = dfa.MakeDFA(
	[][]string{
		{"start", "a", "accept"},
		{"start", "b", "accept"},
		{"start", "c", "accept"},

		{"accept", "a", "accept"},
		{"accept", "b", "accept"},
		{"accept", "c", "accept"},
	},
	"start",
	[]string{"start", "accept"})

var redLightDFA, _ = dfa.MakeDFA(
	[][]string{
		{"start", "r", "red"},
		{"start", "y", "yellow"},
		{"start", "g", "green"},

		{"red", "r", "error"},
		{"red", "y", "error"},
		{"red", "g", "green"},

		{"yellow", "r", "red"},
		{"yellow", "y", "error"},
		{"yellow", "g", "error"},

		{"green", "r", "error"},
		{"green", "y", "yellow"},
		{"green", "g", "error"},

		{"error", "r", "error"},
		{"error", "y", "error"},
		{"error", "g", "error"},
	},
	"start",
	[]string{"red", "yellow", "green"})

var mod2DFA, _ = dfa.MakeDFA(
	[][]string{
		{"zero", "a", "one"},

		{"one", "a", "zero"},
	},
	"zero",
	[]string{"zero", "one"})

// set up the makeDFATests automatically
func init() {
	makeDFATests = []makeDFAT{
		{[][]string{{"start", "a", "start"}}, "start", []string{}, nil},
	}
}

// set up the startTests automatically
func init() {
	startTests = []startT{
		{emptyDFA, "emptyDFA", "a b c", "{start [a b c]}"},

		{allDFA, "allDFA", "", "{start []}"},
	}
}

// set up the stepTests automatically
func init() {
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration
	var step3 machine.Configuration
	var step4 machine.Configuration
	var step5 machine.Configuration

	start = emptyDFA.Start("a b")
	step1, _ = emptyDFA.Step(start)
	stepTests = append(stepTests, []stepT{
		{emptyDFA, "emptyDFA", start, "{reject [b]}", nil},
		{emptyDFA, "emptyDFA", step1, "{reject []}", nil},
	}...)

	start = redLightDFA.Start("r g y r g y")
	step1, _ = redLightDFA.Step(start)
	step2, _ = redLightDFA.Step(step1)
	step3, _ = redLightDFA.Step(step2)
	step4, _ = redLightDFA.Step(step3)
	step5, _ = redLightDFA.Step(step4)
	stepTests = append(stepTests, []stepT{
		{redLightDFA, "redLightDFA", start, "{red [g y r g y]}", nil},
		{redLightDFA, "redLightDFA", step1, "{green [y r g y]}", nil},
		{redLightDFA, "redLightDFA", step2, "{yellow [r g y]}", nil},
		{redLightDFA, "redLightDFA", step3, "{red [g y]}", nil},
		{redLightDFA, "redLightDFA", step4, "{green [y]}", nil},
		{redLightDFA, "redLightDFA", step5, "{yellow []}", nil},
	}...)

	start = mod2DFA.Start("a a")
	step1, _ = mod2DFA.Step(start)
	stepTests = append(stepTests, []stepT{
		{mod2DFA, "mod2DFA", start, "{one [a]}", nil},
		{mod2DFA, "mod2DFA", step1, "{zero []}", nil},
	}...)

	start = mod2DFA.Start("a a a")
	step1, _ = mod2DFA.Step(start)
	step2, _ = mod2DFA.Step(step1)
	stepTests = append(stepTests, []stepT{
		{mod2DFA, "mod2DFA", start, "{one [a a]}", nil},
		{mod2DFA, "mod2DFA", step1, "{zero [a]}", nil},
		{mod2DFA, "mod2DFA", step2, "{one []}", nil},
	}...)
}

// set up the isAcceptTests automatically
func init() {
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration
	var step3 machine.Configuration

	start = emptyDFA.Start("a b")
	step1, _ = emptyDFA.Step(start)
	step2, _ = emptyDFA.Step(step1)
	isAcceptTests = append(isAcceptTests, []isAcceptT{
		{emptyDFA, "emptyDFA", start, false},
		{emptyDFA, "emptyDFA", step1, false},
		{emptyDFA, "emptyDFA", step2, false},
	}...)

	start = redLightDFA.Start("r g y")
	step1, _ = redLightDFA.Step(start)
	step2, _ = redLightDFA.Step(step1)
	step3, _ = redLightDFA.Step(step2)
	isAcceptTests = append(isAcceptTests,[]isAcceptT{
		{redLightDFA, "redLightDFA", start, false},
		{redLightDFA, "redLightDFA", step1, false},
		{redLightDFA, "redLightDFA", step2, false},
		{redLightDFA, "redLightDFA", step3, true},
	}...)
}

// set up the isRejectTests automatically
func init() {
	var start machine.Configuration
	var step1 machine.Configuration
	var step2 machine.Configuration
	var step3 machine.Configuration

	start = emptyDFA.Start("a b")
	step1, _ = emptyDFA.Step(start)
	step2, _ = emptyDFA.Step(step1)
	isRejectTests = append(isRejectTests, []isRejectT{
		{emptyDFA, "emptyDFA", start, false},
		{emptyDFA, "emptyDFA", step1, false},
		{emptyDFA, "emptyDFA", step2, true},
	}...)

	start = redLightDFA.Start("r g y")
	step1, _ = redLightDFA.Step(start)
	step2, _ = redLightDFA.Step(step1)
	step3, _ = redLightDFA.Step(step2)
	isRejectTests = append(isRejectTests,[]isRejectT{
		{redLightDFA, "redLightDFA", start, false},
		{redLightDFA, "redLightDFA", step1, false},
		{redLightDFA, "redLightDFA", step2, false},
		{redLightDFA, "redLightDFA", step3, false},
	}...)
}
