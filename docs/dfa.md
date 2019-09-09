# Deterministic Finite Automaton

The YAML file for DFAs can be constructed with,

```
start: STATE
accept-states: [STATES]
transitions:
  - TRANSITION
  - TRANSITION
  - TRANSITION
  ...
```

where

```
STATE --> string
STATES --> STATE
       --> STATE, STATES
TRANSITION --> [STATE, SYMBOL, STATE]
SYMBOL --> string
```

## Example

```
# Recognizes the language of strings with "abc" as a substring
# over the alphabet {"a", "b", "c"}

start: seen0
accept-states: [seenABC]
transitions:
  - [seen0, a, seenA]
  - [seen0, b, seen0]
  - [seen0, c, seen0]

  - [seenA, a, seenA]
  - [seenA, b, seenAB]
  - [seenA, c, seenC]

  - [seenAB, a, seenA]
  - [seenAB, b, seen0]
  - [seenAB, c, seenABC]

  - [seenABC, a, seenABC]
  - [seenABC, b, seenABC]
  - [seenABC, c, seenABC]
```

This example recognizes the language of strings with "abc" as a substring.

## Notes

* Each transition **must be** indented.
The indentation **must be** made with spaces, **not** tabs.

* The states and symbols **can be** quoted.
This means if left unquoted, the YAML interpreter treats these has strings automatically.
The exception is for [special characters](https://yaml.org/spec/1.2/spec.html#Characters).
A rule of thumb: if it is constructed with letters and numbers, it is most likely a string.

* There **must be** a single space after ":", "-", and ",".
There **must not be** spaces before these charaters.

* There **can be** blank lines inbetween transitions, as shown above.

* Comments are made with "#", as shown above.
