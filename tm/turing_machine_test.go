package tm_test

import (
	"testing"
	"strings"
	"fmt"

	"github.com/cjcodell1/tint/tm"
)

func toString(conf tm.Config) string {
	return fmt.Sprintf("(%s, %s, %d)", conf.State, strings.Join(conf.Tape, " "), conf.Index)
}

// Let's make some TMs to use for testing
// All are over the language {a, b, c}
var tmEmpty = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"start", "a"}, tm.Output{"reject", "a", tm.Right}},
		{tm.Input{"start", "b"}, tm.Output{"reject", "b", tm.Right}},
		{tm.Input{"start", "c"}, tm.Output{"reject", "c", tm.Right}},
		{tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
	},
	"start",
	"accept",
	"reject"}

var tmAll = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"q0", "a"}, tm.Output{"accept", "a", tm.Right}},
		{tm.Input{"q0", "b"}, tm.Output{"accept", "b", tm.Right}},
		{tm.Input{"q0", "c"}, tm.Output{"accept", "c", tm.Right}},
		{tm.Input{"q0", tm.Blank}, tm.Output{"accept", tm.Blank, tm.Right}},
	},
	"q0",
	"accept",
	"reject"}

var tmAddMarkers = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"place$", "a"}, tm.Output{"placeA", "$", tm.Right}},
		{tm.Input{"place$", "b"}, tm.Output{"placeB", "$", tm.Right}},
		{tm.Input{"place$", "c"}, tm.Output{"placeC", "$", tm.Right}},
		{tm.Input{"place$", tm.Blank}, tm.Output{"placeLast$", "$", tm.Right}},
		{tm.Input{"place$", "$"}, tm.Output{"reject", "$", tm.Right}},

		{tm.Input{"placeA", "a"}, tm.Output{"placeA", "a", tm.Right}},
		{tm.Input{"placeA", "b"}, tm.Output{"placeB", "a", tm.Right}},
		{tm.Input{"placeA", "c"}, tm.Output{"placeC", "a", tm.Right}},
		{tm.Input{"placeA", tm.Blank}, tm.Output{"placeLast$", "a", tm.Right}},
		{tm.Input{"placeA", "$"}, tm.Output{"reject", "$", tm.Right}},

		{tm.Input{"placeB", "a"}, tm.Output{"placeA", "b", tm.Right}},
		{tm.Input{"placeB", "b"}, tm.Output{"placeB", "b", tm.Right}},
		{tm.Input{"placeB", "c"}, tm.Output{"placeC", "b", tm.Right}},
		{tm.Input{"placeB", tm.Blank}, tm.Output{"placeLast$", "b", tm.Right}},
		{tm.Input{"placeB", "$"}, tm.Output{"reject", "$", tm.Right}},

		{tm.Input{"placeC", "a"}, tm.Output{"placeA", "c", tm.Right}},
		{tm.Input{"placeC", "b"}, tm.Output{"placeB", "c", tm.Right}},
		{tm.Input{"placeC", "c"}, tm.Output{"placeC", "c", tm.Right}},
		{tm.Input{"placeC", tm.Blank}, tm.Output{"placeLast$", "c", tm.Right}},
		{tm.Input{"placeC", "$"}, tm.Output{"reject", "$", tm.Right}},

		{tm.Input{"placeLast$", "a"}, tm.Output{"reject", "a", tm.Right}},
		{tm.Input{"placeLast$", "b"}, tm.Output{"reject", "b", tm.Right}},
		{tm.Input{"placeLast$", "c"}, tm.Output{"reject", "c", tm.Right}},
		{tm.Input{"placeLast$", tm.Blank}, tm.Output{"returnToStart", "$", tm.Left}},
		{tm.Input{"placeLast$", "$"}, tm.Output{"reject", "$", tm.Right}},

		{tm.Input{"returnToStart", "a"}, tm.Output{"returnToStart", "a", tm.Left}},
		{tm.Input{"returnToStart", "b"}, tm.Output{"returnToStart", "b", tm.Left}},
		{tm.Input{"returnToStart", "c"}, tm.Output{"returnToStart", "c", tm.Left}},
		{tm.Input{"returnToStart", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Left}},
		{tm.Input{"returnToStart", "$"}, tm.Output{"accept", "$", tm.Left}},
	},
	"place$",
	"accept",
	"reject"}

var tmMoveRight = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"moveRight", "a"}, tm.Output{"moveRight", "a", tm.Right}},
		{tm.Input{"moveRight", "b"}, tm.Output{"moveRight", "b", tm.Right}},
		{tm.Input{"moveRight", "c"}, tm.Output{"moveRight", "c", tm.Right}},
		{tm.Input{"moveRight", tm.Blank}, tm.Output{"moveRight", tm.Blank, tm.Right}},
	},
	"moveRight",
	"accept",
	"reject"}

var tmMoveLeft = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"moveLeft", "a"}, tm.Output{"moveLeft", "a", tm.Left}},
		{tm.Input{"moveLeft", "b"}, tm.Output{"moveLeft", "b", tm.Left}},
		{tm.Input{"moveLeft", "c"}, tm.Output{"moveLeft", "c", tm.Left}},
		{tm.Input{"moveLeft", tm.Blank}, tm.Output{"moveLeft", tm.Blank, tm.Left}},
	},
	"moveLeft",
	"accept",
	"reject"}

// TM over the language {longSymbol}
var tmLongSymbol = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"start", "longSymbol"}, tm.Output{"accept", "longSymbol", tm.Right}},
		{tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
	},
	"start",
	"accept",
	"reject"}

// TM over the language {a, A}
var tmCaseSens = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"start", "a"}, tm.Output{"accept", "b", tm.Right}},
		{tm.Input{"start", "A"}, tm.Output{"reject", "B", tm.Right}},
		{tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
	},
	"start",
	"accept",
	"reject"}

// TM over the language {a}
var tmBlankRight = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"any", "a"}, tm.Output{"any", tm.Blank, tm.Right}},
		{tm.Input{"any", tm.Blank}, tm.Output{"any", tm.Blank, tm.Right}},
	},
	"any",
	"accept",
	"reject"}

var tmBlankLeft = tm.TuringMachine{
	[]tm.Transition{
		{tm.Input{"any", "a"}, tm.Output{"any", tm.Blank, tm.Left}},
		{tm.Input{"any", tm.Blank}, tm.Output{"any", tm.Blank, tm.Left}},
	},
	"any",
	"accept",
	"reject"}

// Testing Start

type StartTest struct {
	tm tm.TuringMachine
	tmName string
	input string
	expect tm.Config
}

var startTests = []StartTest{
	{tmEmpty, "tmEmpty", "", tm.Config{"start", []string{}, 0}},
	{tmEmpty, "tmEmpty", "a", tm.Config{"start", []string{"a"}, 0}},
	{tmEmpty, "tmEmpty", "b", tm.Config{"start", []string{"b"}, 0}},
	{tmEmpty, "tmEmpty", "c", tm.Config{"start", []string{"c"}, 0}},
	{tmLongSymbol, "tmLongSymbol", "longSymbol", tm.Config{"start", []string{"longSymbol"}, 0}},
	{tmEmpty, "tmEmpty", "a b c", tm.Config{"start", []string{"a", "b", "c"}, 0}},
	{tmEmpty, "tmEmpty", "c c c c c a", tm.Config{"start", []string{"c", "c", "c", "c", "c", "a"}, 0}},
	{tmLongSymbol, "tmLongSymbol", "longSymbol longSymbol", tm.Config{"start", []string{"longSymbol", "longSymbol"}, 0}},
	{tmCaseSens, "tmCaseSens", "a A a A", tm.Config{"start", []string{"a", "A", "a", "A"}, 0}},
}

func TestStart(t *testing.T) {
	for _, tc := range startTests {
		got := tc.tm.Start(tc.input)
		if toString(tc.expect) != toString(got) {
			t.Errorf("%s.Start(%s) == %s != %s", tc.tmName, tc.input, toString(got), toString(tc.expect))
		}
	}
}

// Testing Step

type StepTest struct {
	tm tm.TuringMachine
	tmName string
	input tm.Config
	expect tm.Config
}

var stepTests = []StepTest{
	{tmEmpty, "tmEmpty", tm.Config{"start", []string{}, 0}, tm.Config{"reject", []string{}, 0}},
	{tmEmpty, "tmEmpty", tm.Config{"start", []string{"a"}, 0}, tm.Config{"reject", []string{"a"}, 1}},
	{tmEmpty, "tmEmpty", tm.Config{"start", []string{"a", "a"}, 0}, tm.Config{"reject", []string{"a", "a"}, 1}},
	{tmEmpty, "tmEmpty", tm.Config{"reject", []string{"a", "a"}, 1}, tm.Config{"reject", []string{"a", "a"}, 1}},

	{tmAll, "tmAll", tm.Config{"q0", []string{}, 0}, tm.Config{"accept", []string{}, 0}},
	{tmAll, "tmAll", tm.Config{"q0", []string{"c"}, 0}, tm.Config{"accept", []string{"c"}, 1}},
	{tmAll, "tmAll", tm.Config{"q0", []string{"c", "a"}, 0}, tm.Config{"accept", []string{"c", "a"}, 1}},
	{tmAll, "tmAll", tm.Config{"accept", []string{"c", "a"}, 1}, tm.Config{"accept", []string{"c", "a"}, 1}},

	{tmAddMarkers, "tmAddMarkers", tm.Config{"place$", []string{}, 0}, tm.Config{"placeLast$", []string{"$"}, 1}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"placeLast$", []string{"$"}, 1}, tm.Config{"returnToStart", []string{"$", "$"}, 0}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "$"}, 0}, tm.Config{"accept", []string{"$", "$"}, 0}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"place$", []string{"c", "b", "a"}, 0}, tm.Config{"placeC", []string{"$", "b", "a"}, 1}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"placeC", []string{"$", "b", "a"}, 1}, tm.Config{"placeB", []string{"$", "c", "a"}, 2}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"placeB", []string{"$", "c", "a"}, 2}, tm.Config{"placeA", []string{"$", "c", "b"}, 3}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"placeA", []string{"$", "c", "b"}, 3}, tm.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}},
	{tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}, tm.Config{"accept", []string{"$", "c", "b", "a", "$"}, 0}},

	{tmBlankRight, "tmBlankRight", tm.Config{"any", []string{}, 0}, tm.Config{"any", []string{}, 0}},
	{tmBlankRight, "tmBlankRight", tm.Config{"any", []string{"a"}, 0}, tm.Config{"any", []string{}, 0}},
	{tmBlankRight, "tmBlankRight", tm.Config{"any", []string{"a"}, 1}, tm.Config{"any", []string{"a"}, 1}},

	{tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{}, 0}, tm.Config{"any", []string{}, 0}},
	{tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{"a"}, 0}, tm.Config{"any", []string{}, 0}},
	{tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{"a"}, 1}, tm.Config{"any", []string{"a"}, 0}},

	{tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 0}, tm.Config{"moveRight", []string{"a", "b", "c"}, 1}},
	{tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 1}, tm.Config{"moveRight", []string{"a", "b", "c"}, 2}},
	{tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 2}, tm.Config{"moveRight", []string{"a", "b", "c"}, 3}},
	{tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 3}, tm.Config{"moveRight", []string{"a", "b", "c"}, 3}},

	{tmMoveLeft, "tmMoveLeft", tm.Config{"moveLeft", []string{"a", "b", "c"}, 0}, tm.Config{"moveLeft", []string{"a", "b", "c"}, 0}},

	{tmCaseSens, "tmCaseSens", tm.Config{"start", []string{"a"}, 0}, tm.Config{"accept", []string{"b"}, 1}},
	{tmCaseSens, "tmCaseSens", tm.Config{"start", []string{"A"}, 0}, tm.Config{"reject", []string{"B"}, 1}},
}

func TestStep(t *testing.T) {
	for _, tc := range stepTests {
		got := tc.tm.Step(tc.input)
		if toString(tc.expect) != toString(got) {
			t.Errorf("%s.Step(%s) == %s != %s", tc.tmName, toString(tc.input), toString(got), toString(tc.expect))
		}
	}
}
