# tint - A Turing machine Interpreter

`tint` can interpret Turing machines and other models of computation like DFAs, NFAs, and PDAs.
This allows one to program their own model of computation and simulate it with inputs.
See a machine's documentation for specifics on how to program it.

## Using tint

```
./tint -m MACHINE_TYPE MACHINE_FILE TEST_FILE
```

To use `tint` you must use the **-m** flag to specify a machine type.
Current and future machine include:
- "dfa"
- "nfa" (planned)
- "pda" (planned)
- "tm"

The machine file is a YAML-specified machine with listed states and transitions.
See each machine's documentation on how to format this file.

The test file is used to simulate the machine.
On each line, there is a single, unquoted test.
For example,
```
a a b b a a
a b a

one two three
```
is an example of a test file with four different tests.
The first two tests are "a a b b a a" and "a b a".
The spaces separate each symbol, so there are six symbols in the first test and three in the second.
The third test is the empty string.
The empty string is denoted with a blank line.
Be **careful** about leaving a blank line at the end of your file, you might unexpectedly test the empty string.
The final test shows that symbols can be mutliple characters long; each symbol is separated with a space.

The last two flags are the **-v** and **-t** flags.
The **-v** flag prints each simulation verbosely: step by step.
The **-t** flag interprets the test file as a single, quoted test.
This is helpful for quickly testing a machine has it is being built.
Here is an example of using this flag:
> ./tint -m dfa -t my_dfa.yaml "a b c"

## Common Mistakes

* Leaving out indentation for the transitions.
* Using tabs instead of spaces for transitions.
* Leaving a [special character](https://yaml.org/spec/1.2/spec.html#id2772075) unquoted.
* Not putting a space after ":", "-", or ",".
* Not putting a "#" to begin a comment.
* Forgetting to put commas (",") inbetween states or values in a transition.
* Misspelling.
* Copy and paste errors.
