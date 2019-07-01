# tint

tint is a Turing machine interpreter written in Go.
It is a great environment to design and easily test Turing machines.

## Goal

My goal for tint is that it gives people learning about Turing machines some hands-on experience.
When I learned about finite state automata as an undergraduate, my favorite part was building the machines and understanding how each one works.

## Install

First [download](https://golang.org/dl/) and then [configure](https://golang.org/doc/install) Go.

Then download and install tint:

`go get github.com/cjcodell1/tint`

If you've set up Go correctly, you should be able to enter `tint` and see the help message.

## Features

- [X] Simulate Turing machines with test(s).
- [X] Build Turing machines from YAML files.
- [X] CLI to operate the program.

## Plans

- [ ] Prettier printing in verbose mode.
- [ ] Stepper to step forward through a test.
- [ ] Implement other machines: DFA, NFA, PDA, multi-tape TM, etc..
