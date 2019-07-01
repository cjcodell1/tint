# tint

tint is a Turing machine interpreter written in Go.
It is a great environment to design and easily test Turing machines.

## Goal

My goal for tint is that it gives people learning about Turing machines some hands-on experience.
When I learned about finite state automata as an undergraduate, my favorite part was building the machines and understanding how each one works.

## Install

First [download](https://golang.org/dl/) and then [configure](https://golang.org/doc/install) Go. Then download and install tint:

`go get -u github.com/cjcodell1/tint`

If you've set up Go correctly, you should be able to enter `tint` and see the help message.

## Use

To use this program you will need a Turing machine and test(s).

### Building the Turing Machine

Use YAML to specify the Turing machine. The below code shows an example of a YAML specification:

```yaml
# file: example.yaml
# recognizes the language aa*bb*aa*
# over the alphabet {a, b}

start: seen0
accept: good
reject: bad
transitions:
    # find the first run of 'a'
    - [seen0, a, seenA, c, R]
    - [seen0, '*', bad, c, R]

    # find the first run of 'b'
    - [seenA, a, seenA, c, R]
    - [seenA, b, seenAB, c, R]
    - [seenA, _, bad, c, R]

    # find the next run of 'a'
    - [seenAB, a, seenABA, c, R]
    - [seenAB, b, seenAB, c, R]
    - [seenAB, _, bad, c, R]

    # expect only 'a', accept on '_'
    - [seenABA, a, seenABA, c, R]
    - [seenABA, b, bad, c, R]
    - [seenABA, _, good, c, R]
```

YAML is pretty similar to JSON; it supports maps of key-value pairs, lists, strings, and comments ([among others](https://yaml.org/)).
One **important note** about YAML is that indentations must be made with **spaces**, **not tabs**.
Using tabs to indent lines will cause an error.
Every Turing machine has four keys: `start`, `accept`, `reject`, and `transitions`.
The order of these keys does not matter, but for readability it is best to keep the transitions at the botton.

`start`, `accept`, and `reject`are respectively the start state, accept state, and reject state.
In the example above the start state is called `seen0`, the accept state is called `good`, and the reject state is called `bad`.
These states can be named anything as long as the accept state and reject state are named differently.

`transitions` specify a list of transitions for the Turing machine.
Each transition is of the form
![Transition Function](https://latex.codecogs.com/gif.latex?\delta:&space;Q&space;\times&space;\Gamma&space;\to&space;Q&space;\times&space;\Gamma&space;\times&space;\{\text{L},&space;\text{R}\}), where ![Q](https://latex.codecogs.com/gif.latex?Q) is the set of states and ![Gamma](https://latex.codecogs.com/gif.latex?\Gamma) is the tape alphabet.

Basically a transition is `[current_state, read_symbol, next_state, write_symbol, move_head]`.
In YAML, strings do not always have to be placed in double or single quotes.
However, if your states or symbols are special, non-alphanumeric characters then you may need to use quotes to denote a string (e.g. "\*", "$", " ").

For the reasons below, you **cannot** use an asterisk "\*" as a state, symbol in the input alphabet, or symbol in the tape alphabet.
You **cannot** use an underscore "\_" as a state or symbol in the input alphabet.
You **cannot** use a space " " as a symbol in the input alphabet and **should not** use it as a state or symbol in the tape alphabet.

"\*" is a special character which denotes "any other".
For example "\*" is used as the `read_symbol` on line 11 of the above example.
This tells the Turing machine if it is in state `seen0` reads any symbol other than `a` (symbols listed above the "\*" line), then perform that transition on line 11.
"\*" can also be used as the `write_symbol`, telling the Turing machine to write the same symbol it had just read.
Finally, "\*" can be used as `current_state` and `next_state` giving similar effects.
One **important note** on the use of "\*" is this interpreter will use the first matching transition for a state, symbol pair.
So, swapping lines 10 and 11 of the above file will cause the Turing machine to recognize the empty language.

"\_" is also a special character which denotes the blank symbol.

A space " " does not have any special meaning in transitions like the above two symbols, however a configuration's tape is prined with spaces seperating the symbols, so it would be confusing to use as a state or symbol.
Futhermore it is used as a delimiter for the input string, so it cannot be used as a symbol in the input string.

### Input Test(s)

There are two ways to simulate the Turing machine with inputs: with a single test (see the discussion of flags) or with a file of tests.
An example testing file is:

```
a b a
a a a b a a
b

a b
```

Each test is **unquoted** and on **separate lines** (DOS or Unix line-endings).
Each symbol in a test is delimited with one or more spaces.
To test the empty string, just use a blank line (e.g. line 4).

### tint

Assuming the above YAML Turing machine specification is called `example.yaml` and the above testing file is called `tests.txt`, here are some examples of my program.

```
$ tint examples.yaml tests.txt
Simulating with "a b a".
Accepted.

Simulating with "".
Rejected.

Simulating with "a b".
Rejected.

Simulating with "a a a b a a".
Accepted.

Simulating with "b".
Rejected.

```
```
$ tint -v examples.yaml tests.txt
Simulating with "a a a b a a".
seen0: "a a a b a a", at 0
seenA: "c a a b a a", at 1
seenA: "c c a b a a", at 2
seenA: "c c c b a a", at 3
seenAB: "c c c c a a", at 4
seenABA: "c c c c c a", at 5
seenABA: "c c c c c c", at 6
good: "c c c c c c c", at 7
Accepted.

Simulating with "b".
seen0: "b", at 0
bad: "c", at 1
Rejected.

Simulating with "".
seen0: "", at 0
bad: "c", at 1
Rejected.

Simulating with "a b a".
seen0: "a b a", at 0
seenA: "c b a", at 1
seenAB: "c c a", at 2
seenABA: "c c c", at 3
good: "c c c c", at 4
Accepted.

Simulating with "a b".
seen0: "a b", at 0
seenA: "c b", at 1
seenAB: "c c", at 2
bad: "c c c", at 3
Rejected.

```
```
$ tint -t example.yaml "a a b b a a"
Simulating with "a a b b a a".
Accepted.

```
The inputs are simulated and printed unordered.

#### Flags

- The `-v, --verbose` flag prints the step-by-step transitions.
- The `-t, --test` flag replaces the testing file with a single, quoted test.

## Features

- [X] Simulate Turing machines with test(s).
- [X] Build Turing machines from YAML files.
- [X] CLI to operate the program.

## Plans

- [ ] Prettier printing in verbose mode.
- [ ] Stepper to step forward through a test.
- [ ] Implement other machines: DFA, NFA, PDA, multi-tape TM, etc..
