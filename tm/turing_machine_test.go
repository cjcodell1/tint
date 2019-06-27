package tm_test

import (
    "testing"
    "strings"
    "fmt"

    "github.com/cjcodell1/tint/tm"
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
    trans []tm.Transition
    start string
    accept string
    reject string
    isErrNil bool
}

var newTMTests = []newTMTest{
    {[]tm.Transition{}, "start", "accept", "reject", true},
    {[]tm.Transition{}, "same", "same", "reject", true},
    {[]tm.Transition{}, "same", "accept", "same", true},
    {[]tm.Transition{}, "start", "same", "same", false},
    {[]tm.Transition{}, "same", "same", "same", false},
}

func testNewTuringMachine(t *testing.T) {
    for _, tc := range newTMTests {
        got, gotErr := tm.NewTuringMachine(tc.trans, tc.start, tc.accept, tc.reject)

        var errString string
        if tc.isErrNil {
            errString = "nil"
        } else {
            errString = "non-nil"
        }

        if tc.isErrNil && (gotErr != nil) {
            if gotErr == nil{
                t.Errorf("NewTuringMachine(%v, %s, %s, %s) = %v, %s != someTM, %s", tc.trans, tc.start, tc.accept, tc.reject, got, "nil", errString)
            } else {
                t.Errorf("NewTuringMachine(%v, %s, %s, %s) = %v, %s != someTM, %s", tc.trans, tc.start, tc.accept, tc.reject, got, gotErr.Error(), errString)
            }
        }
    }
}

// Testing Start

type startTest struct {
    tm tm.TuringMachine
    tmName string
    input string
    expect tm.Config
}

var startTests = []startTest {
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
    tm tm.TuringMachine
    tmName string
    input tm.Config
    expect tm.Config
    isErrNil bool
}

var stepTests = []stepTest{
    {tmEmpty, "tmEmpty", tm.Config{"start", []string{}, 0}, tm.Config{"reject", []string{}, 0}, true},
    {tmEmpty, "tmEmpty", tm.Config{"start", []string{"a"}, 0}, tm.Config{"reject", []string{"a"}, 1}, true},
    {tmEmpty, "tmEmpty", tm.Config{"start", []string{"a", "a"}, 0}, tm.Config{"reject", []string{"a", "a"}, 1}, true},
    {tmEmpty, "tmEmpty", tm.Config{"reject", []string{"a", "a"}, 1}, tm.Config{"reject", []string{"a", "a"}, 1}, true},

    {tmAll, "tmAll", tm.Config{"q0", []string{}, 0}, tm.Config{"accept", []string{}, 0}, true},
    {tmAll, "tmAll", tm.Config{"q0", []string{"c"}, 0}, tm.Config{"accept", []string{"c"}, 1}, true},
    {tmAll, "tmAll", tm.Config{"q0", []string{"c", "a"}, 0}, tm.Config{"accept", []string{"c", "a"}, 1}, true},
    {tmAll, "tmAll", tm.Config{"accept", []string{"c", "a"}, 1}, tm.Config{"accept", []string{"c", "a"}, 1}, true},

    {tmAddMarkers, "tmAddMarkers", tm.Config{"place$", []string{}, 0}, tm.Config{"placeLast$", []string{"$"}, 1}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"placeLast$", []string{"$"}, 1}, tm.Config{"returnToStart", []string{"$", "$"}, 0}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "$"}, 0}, tm.Config{"all done", []string{"$", "$"}, 0}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"place$", []string{"c", "b", "a"}, 0}, tm.Config{"placeC", []string{"$", "b", "a"}, 1}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"placeC", []string{"$", "b", "a"}, 1}, tm.Config{"placeB", []string{"$", "c", "a"}, 2}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"placeB", []string{"$", "c", "a"}, 2}, tm.Config{"placeA", []string{"$", "c", "b"}, 3}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"placeA", []string{"$", "c", "b"}, 3}, tm.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"placeLast$", []string{"$", "c", "b", "a"}, 4}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 3}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 2}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 1}, tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"returnToStart", []string{"$", "c", "b", "a", "$"}, 0}, tm.Config{"all done", []string{"$", "c", "b", "a", "$"}, 0}, true},

    {tmBlankRight, "tmBlankRight", tm.Config{"any", []string{}, 0}, tm.Config{"any", []string{}, 0}, true},
    {tmBlankRight, "tmBlankRight", tm.Config{"any", []string{"a"}, 0}, tm.Config{"any", []string{}, 0}, true},
    {tmBlankRight, "tmBlankRight", tm.Config{"any", []string{"a"}, 1}, tm.Config{"any", []string{"a"}, 1}, true},

    {tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{}, 0}, tm.Config{"any", []string{}, 0}, true},
    {tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{"a"}, 0}, tm.Config{"any", []string{}, 0}, true},
    {tmBlankLeft, "tmBlankLeft", tm.Config{"any", []string{"a"}, 1}, tm.Config{"any", []string{"a"}, 0}, true},

    {tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 0}, tm.Config{"moveRight", []string{"a", "b", "c"}, 1}, true},
    {tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 1}, tm.Config{"moveRight", []string{"a", "b", "c"}, 2}, true},
    {tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 2}, tm.Config{"moveRight", []string{"a", "b", "c"}, 3}, true},
    {tmMoveRight, "tmMoveRight", tm.Config{"moveRight", []string{"a", "b", "c"}, 3}, tm.Config{"moveRight", []string{"a", "b", "c"}, 3}, true},

    {tmMoveLeft, "tmMoveLeft", tm.Config{"moveLeft", []string{"a", "b", "c"}, 0}, tm.Config{"moveLeft", []string{"a", "b", "c"}, 0}, true},

    {tmCaseSens, "tmCaseSens", tm.Config{"start", []string{"a"}, 0}, tm.Config{"accept", []string{"b"}, 1}, true},
    {tmCaseSens, "tmCaseSens", tm.Config{"start", []string{"A"}, 0}, tm.Config{"reject", []string{"B"}, 1}, true},

    {tmWild, "tmWild", tm.Config{"start", []string{"a"}, 0}, tm.Config{"q2", []string{"x"}, 1}, true},
    {tmWild, "tmWild", tm.Config{"q2", []string{"x"}, 1}, tm.Config{"reject", []string{"x", "x"}, 2}, true},
    {tmWild, "tmWild", tm.Config{"start", []string{"a", "b"}, 0}, tm.Config{"q2", []string{"x", "b"}, 1}, true},
    {tmWild, "tmWild", tm.Config{"q2", []string{"x", "b"}, 1}, tm.Config{"accept", []string{"x", "x"}, 2}, true},
    {tmWild, "tmWild", tm.Config{"start", []string{"b", "b"}, 0}, tm.Config{"reject", []string{"x", "b"}, 1}, true},

    {tmWildFirst, "tmWildFirst", tm.Config{"start", []string{"a"}, 0}, tm.Config{"q2", []string{"x"}, 1}, true},
    {tmWildFirst, "tmWildFirst", tm.Config{"start", []string{"b", "b"}, 0}, tm.Config{"q2", []string{"x", "b"}, 1}, true},
    {tmWildFirst, "tmWildFirst", tm.Config{"q2", []string{"x", "b"}, 1}, tm.Config{"reject", []string{"x", "x"}, 2}, true},
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
    tm tm.TuringMachine
    tmName string
    conf tm.Config
    expect bool
}

var isAcceptTests = []isAcceptRejectTest{
    {tmEmpty, "tmEmpty", tm.Config{"accept", []string{}, 0}, true},
    {tmAll, "tmAll", tm.Config{"accept", []string{}, 0}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"all done", []string{}, 0}, true},

    {tmEmpty, "tmEmpty", tm.Config{"should be false", []string{}, 0}, false},
    {tmEmpty, "tmEmpty", tm.Config{"reject", []string{}, 0}, false},
    {tmAll, "tmAll", tm.Config{"Accept", []string{}, 0}, false},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"accept", []string{}, 0}, false},
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
    {tmEmpty, "tmEmpty", tm.Config{"reject", []string{}, 0}, true},
    {tmAll, "tmAll", tm.Config{"reject", []string{}, 0}, true},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"doh", []string{}, 0}, true},

    {tmEmpty, "tmEmpty", tm.Config{"should be false", []string{}, 0}, false},
    {tmEmpty, "tmEmpty", tm.Config{"accept", []string{}, 0}, false},
    {tmAll, "tmAll", tm.Config{"Reject", []string{}, 0}, false},
    {tmAddMarkers, "tmAddMarkers", tm.Config{"accept", []string{}, 0}, false},
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
var tmEmpty, errEmpty = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"start", "a"}, tm.Output{"reject", "a", tm.Right}},
        {tm.Input{"start", "b"}, tm.Output{"reject", "b", tm.Right}},
        {tm.Input{"start", "c"}, tm.Output{"reject", "c", tm.Right}},
        {tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
    },
    "start",
    "accept",
    "reject")

var tmAll, errAll = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"q0", "a"}, tm.Output{"accept", "a", tm.Right}},
        {tm.Input{"q0", "b"}, tm.Output{"accept", "b", tm.Right}},
        {tm.Input{"q0", "c"}, tm.Output{"accept", "c", tm.Right}},
        {tm.Input{"q0", tm.Blank}, tm.Output{"accept", tm.Blank, tm.Right}},
    },
    "q0",
    "accept",
    "reject")

var tmAddMarkers, errAddMarkers = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"place$", "a"}, tm.Output{"placeA", "$", tm.Right}},
        {tm.Input{"place$", "b"}, tm.Output{"placeB", "$", tm.Right}},
        {tm.Input{"place$", "c"}, tm.Output{"placeC", "$", tm.Right}},
        {tm.Input{"place$", tm.Blank}, tm.Output{"placeLast$", "$", tm.Right}},
        {tm.Input{"place$", "$"}, tm.Output{"doh", "$", tm.Right}},

        {tm.Input{"placeA", "a"}, tm.Output{"placeA", "a", tm.Right}},
        {tm.Input{"placeA", "b"}, tm.Output{"placeB", "a", tm.Right}},
        {tm.Input{"placeA", "c"}, tm.Output{"placeC", "a", tm.Right}},
        {tm.Input{"placeA", tm.Blank}, tm.Output{"placeLast$", "a", tm.Right}},
        {tm.Input{"placeA", "$"}, tm.Output{"doh", "$", tm.Right}},

        {tm.Input{"placeB", "a"}, tm.Output{"placeA", "b", tm.Right}},
        {tm.Input{"placeB", "b"}, tm.Output{"placeB", "b", tm.Right}},
        {tm.Input{"placeB", "c"}, tm.Output{"placeC", "b", tm.Right}},
        {tm.Input{"placeB", tm.Blank}, tm.Output{"placeLast$", "b", tm.Right}},
        {tm.Input{"placeB", "$"}, tm.Output{"doh", "$", tm.Right}},

        {tm.Input{"placeC", "a"}, tm.Output{"placeA", "c", tm.Right}},
        {tm.Input{"placeC", "b"}, tm.Output{"placeB", "c", tm.Right}},
        {tm.Input{"placeC", "c"}, tm.Output{"placeC", "c", tm.Right}},
        {tm.Input{"placeC", tm.Blank}, tm.Output{"placeLast$", "c", tm.Right}},
        {tm.Input{"placeC", "$"}, tm.Output{"doh", "$", tm.Right}},

        {tm.Input{"placeLast$", "a"}, tm.Output{"doh", "a", tm.Right}},
        {tm.Input{"placeLast$", "b"}, tm.Output{"doh", "b", tm.Right}},
        {tm.Input{"placeLast$", "c"}, tm.Output{"doh", "c", tm.Right}},
        {tm.Input{"placeLast$", tm.Blank}, tm.Output{"returnToStart", "$", tm.Left}},
        {tm.Input{"placeLast$", "$"}, tm.Output{"doh", "$", tm.Right}},

        {tm.Input{"returnToStart", "a"}, tm.Output{"returnToStart", "a", tm.Left}},
        {tm.Input{"returnToStart", "b"}, tm.Output{"returnToStart", "b", tm.Left}},
        {tm.Input{"returnToStart", "c"}, tm.Output{"returnToStart", "c", tm.Left}},
        {tm.Input{"returnToStart", tm.Blank}, tm.Output{"doh", tm.Blank, tm.Left}},
        {tm.Input{"returnToStart", "$"}, tm.Output{"all done", "$", tm.Left}},
    },
    "place$",
    "all done",
    "doh")

var tmMoveRight, errMoveRight = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"moveRight", "a"}, tm.Output{"moveRight", "a", tm.Right}},
        {tm.Input{"moveRight", "b"}, tm.Output{"moveRight", "b", tm.Right}},
        {tm.Input{"moveRight", "c"}, tm.Output{"moveRight", "c", tm.Right}},
        {tm.Input{"moveRight", tm.Blank}, tm.Output{"moveRight", tm.Blank, tm.Right}},
    },
    "moveRight",
    "accept",
    "reject")

var tmMoveLeft, errMoveLeft= tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"moveLeft", "a"}, tm.Output{"moveLeft", "a", tm.Left}},
        {tm.Input{"moveLeft", "b"}, tm.Output{"moveLeft", "b", tm.Left}},
        {tm.Input{"moveLeft", "c"}, tm.Output{"moveLeft", "c", tm.Left}},
        {tm.Input{"moveLeft", tm.Blank}, tm.Output{"moveLeft", tm.Blank, tm.Left}},
    },
    "moveLeft",
    "accept",
    "reject")

// TM over the language {longSymbol}
var tmLongSymbol, errLongSymbol = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"start", "longSymbol"}, tm.Output{"accept", "longSymbol", tm.Right}},
        {tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
    },
    "start",
    "accept",
    "reject")

// TM over the language {a, A}
var tmCaseSens, errCaseSens = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"start", "a"}, tm.Output{"accept", "b", tm.Right}},
        {tm.Input{"start", "A"}, tm.Output{"reject", "B", tm.Right}},
        {tm.Input{"start", tm.Blank}, tm.Output{"reject", tm.Blank, tm.Right}},
    },
    "start",
    "accept",
    "reject")

// TM over the language {a}
var tmBlankRight, errBlankRight  = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"any", "a"}, tm.Output{"any", tm.Blank, tm.Right}},
        {tm.Input{"any", tm.Blank}, tm.Output{"any", tm.Blank, tm.Right}},
    },
    "any",
    "accept",
    "reject")

var tmBlankLeft, errBlankLeft = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"any", "a"}, tm.Output{"any", tm.Blank, tm.Left}},
        {tm.Input{"any", tm.Blank}, tm.Output{"any", tm.Blank, tm.Left}},
    },
    "any",
    "accept",
    "reject")

var tmWild, errWild = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"start", "a"}, tm.Output{"q2", "x", tm.Right}},
        {tm.Input{"start", "*"}, tm.Output{"reject", "x", tm.Right}},
        {tm.Input{"*", "b"}, tm.Output{"accept", "x", tm.Right}},
        {tm.Input{"*", "*"}, tm.Output{"reject", "x", tm.Right}},
    },
    "start",
    "accept",
    "reject")

var tmWildFirst, errWildFirst = tm.NewTuringMachine(
    []tm.Transition{
        {tm.Input{"start", "*"}, tm.Output{"q2", "x", tm.Right}},
        {tm.Input{"start", "a"}, tm.Output{"reject", "x", tm.Right}},
        {tm.Input{"*", "*"}, tm.Output{"reject", "x", tm.Right}},
        {tm.Input{"*", "b"}, tm.Output{"accept", "x", tm.Right}},
    },
    "start",
    "accept",
    "reject")

func toString(conf tm.Config) string {
    return fmt.Sprintf("(%s, %s, %d)", conf.State, strings.Join(conf.Tape, " "), conf.Index)
}
