# Turing Machine

## Usage

```
./tint -m one-way-tm my_tm1.yaml my_tests.txt
```
```
./tint -m two-way-tm -v my_tm2.yaml my_reject_tests.txt
```
```
./tint -m two-way-tm -v -t my_tm3.yaml "this should accept"
```

## Formal Grammar

The YAML file for TMs can be constructed with,

```
start: STATE
accept: STATE
reject: STATE
transitions:
  - TRANSITION
  - TRANSITION
  - TRANSITION
  ...
```

where

```
STATE --> string
TRANSITION --> [STATE, SYMBOL, STATE, SYMBOL, DIRECTION]
SYMBOL --> string
DIRECTION --> "L"
          --> "R"
```




## Example

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

This example recognizes the language of strings in the form aa\*bb\*aa\*.

## Notes

Every Turing machine has four keys: `start`, `accept`, `reject`, and `transitions`.
The order of these keys does not matter, but for readability it is best to keep the transitions at the botton.

`start`, `accept`, and `reject`are respectively the start state, accept state, and reject state.
In the example above the start state is called `seen0`, the accept state is called `good`, and the reject state is called `bad`.
These states can be named anything as long as the accept state and reject state are named differently.

`transitions` specify a list of transitions for the Turing machine.
Each transition is of the form
![Transition Function](https://latex.codecogs.com/gif.latex?\delta:&space;Q&space;\times&space;\Gamma&space;\to&space;Q&space;\times&space;\Gamma&space;\times&space;\{\text{L},&space;\text{R}\}), where ![Q](https://latex.codecogs.com/gif.latex?Q) is the set of states, ![Gamma](https://latex.codecogs.com/gif.latex?\Gamma) is the tape alphabet, ![L](https://latex.codecogs.com/gif.latex?L) and ![R](https://latex.codecogs.com/gif.latex?R) are the right or left direction for moving the head.

Basically a transition is `[current_state, read_symbol, next_state, write_symbol, move_head]`.
In YAML, strings do not always have to be placed in double or single quotes.
However, if your states or symbols are special, non-alphanumeric characters then you may need to use quotes to denote a string (e.g. "\*", "$", " ").

For the reasons below:
- You **cannot** use an asterisk "\*" as a state, symbol in the input alphabet, or symbol in the tape alphabet.
- You **cannot** use an underscore "\_" as a state or symbol in the input alphabet.
- You **cannot** use a space " " as a symbol in the input alphabet and **should not** use it as a state or symbol in the tape alphabet.

\* is a special character which denotes "any other".
For example \* is used as the `read_symbol` on line 11 of the above example.
This tells the Turing machine if it is in state `seen0` and reads any symbol other than `a` (symbols listed above the \* line), then perform that transition on line 11.
\* can also be used as the `write_symbol`, telling the Turing machine to write the same symbol it had just read.
Finally, \* can be used as `current_state` and `next_state` giving similar effects.
One **important note** on the use of \* is that this interpreter will use the first matching transition for a state-symbol pair.
So, swapping lines 10 and 11 of the above file will cause the Turing machine to recognize the empty language.

"\_" is also a special character which denotes the blank symbol.

A space " " does not have any special meaning in transitions like the above two symbols, however a configuration's tape is prined with spaces seperating the symbols, so it would be confusing to use as a state or symbol.
Futhermore it is used as a delimiter for the input string, so it cannot be used as a symbol in the input string.

-----

* Each transition **must be** indented.
The indentation **must be** made with spaces, **not** tabs.

* The direction does not have to be quoted.
I used quotes to differentiate from a symbol in the formal grammar.

* The states, symbols, and directions **can be** quoted.
This means if left unquoted, the YAML interpreter treats these has strings automatically.
The exception is for [special characters](https://yaml.org/spec/1.2/spec.html#id2772075).
A rule of thumb: if it is constructed with letters and numbers, it is most likely a string.

* There **must be** a single space after ":", "-", and ",".
There **must not be** spaces before these charaters.

* There **can be** blank lines inbetween transitions, as shown above.

* Comments are made with "#", as shown above.
