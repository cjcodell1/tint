# tint

tint is a Turing machine interpreter written in Go.
It is a great environment to design and easily test Turing machines.

## Goal

My goal for tint is that it gives people learning about Turing machines some hands-on experience.
When I learned about finite state automata as an undergraduate, my favorite part was building the machines and understanding how each one works.

## Plans

- [x] Configuration files to build Turing machines.
    - Can build a Turing machine from a YAML file.
    - Documentation to come with the CLI.
- [ ] CLI to build and test Turing machines.
    - Supply a YAML Turing machine specification and a single test or a file containing multiple tests.
- [ ] QoL changes.
    - Use \* to specify "all other" symbols, instead of having to list every possible symbol.
    - Stepper to step forward through a test.
