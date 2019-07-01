package turing_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cjcodell1/tint/machine/turing"
)

// Run all testing functions
func TestAll(t *testing.T) {
	testNewTuringMachine(t)
	testStart(t)
	testStep(t)
	testIsAccept(t)
	testIsReject(t)
}

// Testing NewTuringMachine

type newTMTest struct {
	trans    []turing.Transition
	start    string
	accept   string
	reject   string
	isErrNil bool
}

var newTMTests = []newTMTest{
	{[]turing.Transition{}, "start", "accept", "reject", true},
	{[]turing.Transition{}, "same", "same", "reject", true},
	{[]turing.Transition{}, "same", "accept", "same", true},
	{[]turing.Transition{}, "start", "same", "same", false},
	{[]turing.Transition{}, "same", "same", "same", false},
}

func testNewTuringMachine(t *testing.T) {
	for _, tc := range newTMTests {
		got, gotErr := turing.NewTuringMachine(tc.trans, tc.start, tc.accept, tc.reject)

		var errString string
		if tc.isErrNil {
			errString = "nil"
		} else {
			errString = "non-nil"
		}

		if tc.isErrNil && (gotErr != nil) {
			if gotErr == nil {
				t.Errorf("NewTuringMachine(%v, %s, %s, %s) = %v, %s != someTM, %s", tc.trans, tc.start, tc.accept, tc.reject, got, "nil", errString)
			} else {
				t.Errorf("NewTuringMachine(%v, %s, %s, %s) = %v, %s != someTM, %s", tc.trans, tc.start, tc.accept, tc.reject, got, gotErr.Error(), errString)
			}
		}
	}
}

// Testing Start

type startTest struct {
	tm     turing.TuringMachine
	tmName string
	input  string
	expect turing.Config
}

var startTests = []startTest{
	{tmEmpty, "tmEmpty", "", turing.Config{"start", []string{}, 0}},
	{tmEmpty, "tmEmpty", "a", turing.Config{"start", []string{"a"}, 0}},
	{tmEmpty, "tmEmpty", "b", turing.Config{"start", []string{"b"}, 0}},
	{tmEmpty, "tmEmpty", "c", turing.Config{"start", []string{"c"}, 0}},
	{tmLongSymbol, "tmLongSymbol", "longSymbol", turing.Config{"start", []string{"longSymbol"}, 0}},
	{tmEmpty, "tmEmpty", "a b c", turing.Config{"start", []string{"a", "b", "c"}, 0}},
	{tmEmpty, "tmEmpty", "c c c c c a", turing.Config{"start", []string{"c", "c", "c", "c", "c", "a"}, 0}},
	{tmLongSymbol, "tmLongSymbol", "longSymbol longSymbol", turing.Config{"start", []string{"longSymbol", "longSymbol"}, 0}},
	{tmCaseSens, "tmCaseSens", "a A a A", turing.Config{"start", []string{"a", "A", "a", "A"}, 0}},
}

func testStart(t *testing.T) {
	for _, tc := range startTests {
		got := tc.tm.Start(tc.input)
		if toString(tc.expect) != toString(got) {
			t.Errorf("%s.Start(%s) == %s != %s", tc.tmName, tc.input, toString(got), toString(tc.expect))
		}
	}
}

// Testing Step

type stepTest struct {
	tm       turing.TuringMachine
	tmName   string
	input    turing.Config
	expect   turing.Config
	isErrNil bool
}

var stepTests = []stepTest{
	{tmEmpty, "tmEmpty", turing.Config{"start", []string{}, 0}, turing.Config{"reject", []string{}, 0}, true},
	{tmEmpty, "tmEmpty", turing.Config{"start", []string{"a"}, 0}, turing.Config{"reject", []string{"a"}, 1}, true},
	{tmEmpty, "tmEmpty", turing.Config{"start", []string{"a", "a"}, 0}, turing.Config{"reject", []string{"a", "a"}, 1}, true},
	{tmEmpty, "tmEmpty", turing.Config{"reject", []string{"a", "a"}, 1}, turing.Config{"reject", []string{"a", "a"}, 1}, true},

	{tmAll, "tmAll", turing.Config{"q0", []string{}, 0}, turing.Config{"accept", []string{}, 0}, true},
	{tmAll, "tmAll", turing.Config{"q0", []string{"c"}, 0}, turing.Config{"accept", []string{"c"}, 1}, true},
	{tmAll, "tmAll", turing.Config{"q0", []string{"c", "a"}, 0}, turing.Config{"accept", []string{"c", "a"}, 1}, true},
	{tmAll, "tmAll", turing.Config{"accept", []string{"c", "a"}, 1}, turing.Config{"accept", []string{"c", "a"}, 1}, true},

	{tmAddMarkers, "tmAddMarkers", turing.Config{"place$", []string{}, 0}, turing.Config{"placeLast$", []string{"$"}, 1}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"placeLast$", []string{"$"}, 1}, turing.Config{"returnToStart", []string{"$", "$"}, 0}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"returnToStart", []string{"$", "$"}, 0}, turing.Config{"all done", []string{"$", "$"}, 0}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"place$", []string{"c", "b", "a"}, 0}, turing.Config{"placeC", []string{"$", "b", "a"}, 1}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"placeC", []string{"$", "b", "a"}, 1}, turing.Config{"placeB", []string{"$", "c", "a"}, 2}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"placeB", []string{"$", "c", "a"}, 2}, turing.Config{"placeA", []string{"$", "c", "b"}, 3}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"placeA", []string{"$", "c", "b"}, 3}, turing.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}, turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}, turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}, turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}, turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}, turing.Config{"all done", []string{"$", "c", "b", "a", "$"}, 0}, true},

	{tmBlankRight, "tmBlankRight", turing.Config{"any", []string{}, 0}, turing.Config{"any", []string{}, 0}, true},
	{tmBlankRight, "tmBlankRight", turing.Config{"any", []string{"a"}, 0}, turing.Config{"any", []string{}, 0}, true},
	{tmBlankRight, "tmBlankRight", turing.Config{"any", []string{"a"}, 1}, turing.Config{"any", []string{"a"}, 1}, true},

	{tmBlankLeft, "tmBlankLeft", turing.Config{"any", []string{}, 0}, turing.Config{"any", []string{}, 0}, true},
	{tmBlankLeft, "tmBlankLeft", turing.Config{"any", []string{"a"}, 0}, turing.Config{"any", []string{}, 0}, true},
	{tmBlankLeft, "tmBlankLeft", turing.Config{"any", []string{"a"}, 1}, turing.Config{"any", []string{"a"}, 0}, true},

	{tmMoveRight, "tmMoveRight", turing.Config{"moveRight", []string{"a", "b", "c"}, 0}, turing.Config{"moveRight", []string{"a", "b", "c"}, 1}, true},
	{tmMoveRight, "tmMoveRight", turing.Config{"moveRight", []string{"a", "b", "c"}, 1}, turing.Config{"moveRight", []string{"a", "b", "c"}, 2}, true},
	{tmMoveRight, "tmMoveRight", turing.Config{"moveRight", []string{"a", "b", "c"}, 2}, turing.Config{"moveRight", []string{"a", "b", "c"}, 3}, true},
	{tmMoveRight, "tmMoveRight", turing.Config{"moveRight", []string{"a", "b", "c"}, 3}, turing.Config{"moveRight", []string{"a", "b", "c"}, 3}, true},

	{tmMoveLeft, "tmMoveLeft", turing.Config{"moveLeft", []string{"a", "b", "c"}, 0}, turing.Config{"moveLeft", []string{"a", "b", "c"}, 0}, true},

	{tmCaseSens, "tmCaseSens", turing.Config{"start", []string{"a"}, 0}, turing.Config{"accept", []string{"b"}, 1}, true},
	{tmCaseSens, "tmCaseSens", turing.Config{"start", []string{"A"}, 0}, turing.Config{"reject", []string{"B"}, 1}, true},

	{tmWild, "tmWild", turing.Config{"start", []string{"a"}, 0}, turing.Config{"q2", []string{"x"}, 1}, true},
	{tmWild, "tmWild", turing.Config{"q2", []string{"x"}, 1}, turing.Config{"reject", []string{"x", "x"}, 2}, true},
	{tmWild, "tmWild", turing.Config{"start", []string{"a", "b"}, 0}, turing.Config{"q2", []string{"x", "b"}, 1}, true},
	{tmWild, "tmWild", turing.Config{"q2", []string{"x", "b"}, 1}, turing.Config{"accept", []string{"x", "x"}, 2}, true},
	{tmWild, "tmWild", turing.Config{"start", []string{"b", "b"}, 0}, turing.Config{"reject", []string{"x", "b"}, 1}, true},

	{tmWildFirst, "tmWildFirst", turing.Config{"start", []string{"a"}, 0}, turing.Config{"q2", []string{"x"}, 1}, true},
	{tmWildFirst, "tmWildFirst", turing.Config{"start", []string{"b", "b"}, 0}, turing.Config{"q2", []string{"x", "b"}, 1}, true},
	{tmWildFirst, "tmWildFirst", turing.Config{"q2", []string{"x", "b"}, 1}, turing.Config{"reject", []string{"x", "x"}, 2}, true},

	{tmWriteSame, "tmWriteSame", turing.Config{"same", []string{"a", "b", "c"}, 0}, turing.Config{"same", []string{"a", "b", "c"}, 1}, true},
	{tmWriteSame, "tmWriteSame", turing.Config{"same", []string{"a", "b", "c"}, 1}, turing.Config{"same", []string{"a", "b", "c"}, 2}, true},
	{tmWriteSame, "tmWriteSame", turing.Config{"same", []string{"a", "b", "c"}, 2}, turing.Config{"same", []string{"a", "b", "c"}, 3}, true},
	{tmWriteSame, "tmWriteSame", turing.Config{"same", []string{"a", "b", "c"}, 3}, turing.Config{"accept", []string{"a", "b", "c"}, 3}, true},
}

func testStep(t *testing.T) {
	for _, tc := range stepTests {
		got, gotErr := tc.tm.Step(tc.input)

		var errString string
		if tc.isErrNil {
			errString = "nil"
		} else {
			errString = "not-nil"
		}

		if (toString(tc.expect) != toString(got)) || (tc.isErrNil && (gotErr != nil)) {
			if gotErr == nil {
				t.Errorf("%s.Step(%s) == %s, %s != %s, %s", tc.tmName, toString(tc.input), toString(got), "nil", toString(tc.expect), errString)
			} else {
				t.Errorf("%s.Step(%s) == %s, %s != %s, %s", tc.tmName, toString(tc.input), toString(got), gotErr.Error(), toString(tc.expect), errString)
			}
		}
	}
}

type isAcceptRejectTest struct {
	tm     turing.TuringMachine
	tmName string
	conf   turing.Config
	expect bool
}

var isAcceptTests = []isAcceptRejectTest{
	{tmEmpty, "tmEmpty", turing.Config{"accept", []string{}, 0}, true},
	{tmAll, "tmAll", turing.Config{"accept", []string{}, 0}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"all done", []string{}, 0}, true},

	{tmEmpty, "tmEmpty", turing.Config{"should be false", []string{}, 0}, false},
	{tmEmpty, "tmEmpty", turing.Config{"reject", []string{}, 0}, false},
	{tmAll, "tmAll", turing.Config{"Accept", []string{}, 0}, false},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"accept", []string{}, 0}, false},
}

func testIsAccept(t *testing.T) {
	for _, tc := range isAcceptTests {
		got := tc.tm.IsAccept(tc.conf)
		if got != tc.expect {
			t.Errorf("%s.IsAccept(%s) == %t != %t", tc.tmName, toString(tc.conf), got, tc.expect)
		}
	}
}

var isRejectTests = []isAcceptRejectTest{
	{tmEmpty, "tmEmpty", turing.Config{"reject", []string{}, 0}, true},
	{tmAll, "tmAll", turing.Config{"reject", []string{}, 0}, true},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"doh", []string{}, 0}, true},

	{tmEmpty, "tmEmpty", turing.Config{"should be false", []string{}, 0}, false},
	{tmEmpty, "tmEmpty", turing.Config{"accept", []string{}, 0}, false},
	{tmAll, "tmAll", turing.Config{"Reject", []string{}, 0}, false},
	{tmAddMarkers, "tmAddMarkers", turing.Config{"accept", []string{}, 0}, false},
}

func testIsReject(t *testing.T) {
	for _, tc := range isRejectTests {
		got := tc.tm.IsReject(tc.conf)
		if got != tc.expect {
			t.Errorf("%s.IsReject(%s) == %t != %t", tc.tmName, toString(tc.conf), got, tc.expect)
		}
	}
}

// Let's make some TMs to use for testing
// All are over the language {a, b, c}
var tmEmpty, errEmpty = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"start", "a"}, turing.Output{"reject", "a", turing.Right}},
		{turing.Input{"start", "b"}, turing.Output{"reject", "b", turing.Right}},
		{turing.Input{"start", "c"}, turing.Output{"reject", "c", turing.Right}},
		{turing.Input{"start", turing.Blank}, turing.Output{"reject", turing.Blank, turing.Right}},
	},
	"start",
	"accept",
	"reject")

var tmAll, errAll = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"q0", "a"}, turing.Output{"accept", "a", turing.Right}},
		{turing.Input{"q0", "b"}, turing.Output{"accept", "b", turing.Right}},
		{turing.Input{"q0", "c"}, turing.Output{"accept", "c", turing.Right}},
		{turing.Input{"q0", turing.Blank}, turing.Output{"accept", turing.Blank, turing.Right}},
	},
	"q0",
	"accept",
	"reject")

var tmAddMarkers, errAddMarkers = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"place$", "a"}, turing.Output{"placeA", "$", turing.Right}},
		{turing.Input{"place$", "b"}, turing.Output{"placeB", "$", turing.Right}},
		{turing.Input{"place$", "c"}, turing.Output{"placeC", "$", turing.Right}},
		{turing.Input{"place$", turing.Blank}, turing.Output{"placeLast$", "$", turing.Right}},
		{turing.Input{"place$", "$"}, turing.Output{"doh", "$", turing.Right}},

		{turing.Input{"placeA", "a"}, turing.Output{"placeA", "a", turing.Right}},
		{turing.Input{"placeA", "b"}, turing.Output{"placeB", "a", turing.Right}},
		{turing.Input{"placeA", "c"}, turing.Output{"placeC", "a", turing.Right}},
		{turing.Input{"placeA", turing.Blank}, turing.Output{"placeLast$", "a", turing.Right}},
		{turing.Input{"placeA", "$"}, turing.Output{"doh", "$", turing.Right}},

		{turing.Input{"placeB", "a"}, turing.Output{"placeA", "b", turing.Right}},
		{turing.Input{"placeB", "b"}, turing.Output{"placeB", "b", turing.Right}},
		{turing.Input{"placeB", "c"}, turing.Output{"placeC", "b", turing.Right}},
		{turing.Input{"placeB", turing.Blank}, turing.Output{"placeLast$", "b", turing.Right}},
		{turing.Input{"placeB", "$"}, turing.Output{"doh", "$", turing.Right}},

		{turing.Input{"placeC", "a"}, turing.Output{"placeA", "c", turing.Right}},
		{turing.Input{"placeC", "b"}, turing.Output{"placeB", "c", turing.Right}},
		{turing.Input{"placeC", "c"}, turing.Output{"placeC", "c", turing.Right}},
		{turing.Input{"placeC", turing.Blank}, turing.Output{"placeLast$", "c", turing.Right}},
		{turing.Input{"placeC", "$"}, turing.Output{"doh", "$", turing.Right}},

		{turing.Input{"placeLast$", "a"}, turing.Output{"doh", "a", turing.Right}},
		{turing.Input{"placeLast$", "b"}, turing.Output{"doh", "b", turing.Right}},
		{turing.Input{"placeLast$", "c"}, turing.Output{"doh", "c", turing.Right}},
		{turing.Input{"placeLast$", turing.Blank}, turing.Output{"returnToStart", "$", turing.Left}},
		{turing.Input{"placeLast$", "$"}, turing.Output{"doh", "$", turing.Right}},

		{turing.Input{"returnToStart", "a"}, turing.Output{"returnToStart", "a", turing.Left}},
		{turing.Input{"returnToStart", "b"}, turing.Output{"returnToStart", "b", turing.Left}},
		{turing.Input{"returnToStart", "c"}, turing.Output{"returnToStart", "c", turing.Left}},
		{turing.Input{"returnToStart", turing.Blank}, turing.Output{"doh", turing.Blank, turing.Left}},
		{turing.Input{"returnToStart", "$"}, turing.Output{"all done", "$", turing.Left}},
	},
	"place$",
	"all done",
	"doh")

var tmMoveRight, errMoveRight = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"moveRight", "a"}, turing.Output{"moveRight", "a", turing.Right}},
		{turing.Input{"moveRight", "b"}, turing.Output{"moveRight", "b", turing.Right}},
		{turing.Input{"moveRight", "c"}, turing.Output{"moveRight", "c", turing.Right}},
		{turing.Input{"moveRight", turing.Blank}, turing.Output{"moveRight", turing.Blank, turing.Right}},
	},
	"moveRight",
	"accept",
	"reject")

var tmMoveLeft, errMoveLeft = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"moveLeft", "a"}, turing.Output{"moveLeft", "a", turing.Left}},
		{turing.Input{"moveLeft", "b"}, turing.Output{"moveLeft", "b", turing.Left}},
		{turing.Input{"moveLeft", "c"}, turing.Output{"moveLeft", "c", turing.Left}},
		{turing.Input{"moveLeft", turing.Blank}, turing.Output{"moveLeft", turing.Blank, turing.Left}},
	},
	"moveLeft",
	"accept",
	"reject")

// TM over the language {longSymbol}
var tmLongSymbol, errLongSymbol = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"start", "longSymbol"}, turing.Output{"accept", "longSymbol", turing.Right}},
		{turing.Input{"start", turing.Blank}, turing.Output{"reject", turing.Blank, turing.Right}},
	},
	"start",
	"accept",
	"reject")

// TM over the language {a, A}
var tmCaseSens, errCaseSens = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"start", "a"}, turing.Output{"accept", "b", turing.Right}},
		{turing.Input{"start", "A"}, turing.Output{"reject", "B", turing.Right}},
		{turing.Input{"start", turing.Blank}, turing.Output{"reject", turing.Blank, turing.Right}},
	},
	"start",
	"accept",
	"reject")

// TM over the language {a}
var tmBlankRight, errBlankRight = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"any", "a"}, turing.Output{"any", turing.Blank, turing.Right}},
		{turing.Input{"any", turing.Blank}, turing.Output{"any", turing.Blank, turing.Right}},
	},
	"any",
	"accept",
	"reject")

var tmBlankLeft, errBlankLeft = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"any", "a"}, turing.Output{"any", turing.Blank, turing.Left}},
		{turing.Input{"any", turing.Blank}, turing.Output{"any", turing.Blank, turing.Left}},
	},
	"any",
	"accept",
	"reject")

var tmWild, errWild = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"start", "a"}, turing.Output{"q2", "x", turing.Right}},
		{turing.Input{"start", "*"}, turing.Output{"reject", "x", turing.Right}},
		{turing.Input{"*", "b"}, turing.Output{"accept", "x", turing.Right}},
		{turing.Input{"*", "*"}, turing.Output{"reject", "x", turing.Right}},
	},
	"start",
	"accept",
	"reject")

var tmWildFirst, errWildFirst = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"start", "*"}, turing.Output{"q2", "x", turing.Right}},
		{turing.Input{"start", "a"}, turing.Output{"reject", "x", turing.Right}},
		{turing.Input{"*", "*"}, turing.Output{"reject", "x", turing.Right}},
		{turing.Input{"*", "b"}, turing.Output{"accept", "x", turing.Right}},
	},
	"start",
	"accept",
	"reject")

var tmWriteSame, errWriteSame = turing.NewTuringMachine(
	[]turing.Transition{
		{turing.Input{"same", turing.Blank}, turing.Output{"accept", turing.Blank, turing.Right}},
		{turing.Input{"same", "*"}, turing.Output{"same", "*", turing.Right}},
	},
	"same",
	"accept",
	"reject")

func toString(conf turing.Config) string {
	return fmt.Sprintf("(%s, %s, %d)", conf.State, strings.Join(conf.Tape, " "), conf.Index)
}
