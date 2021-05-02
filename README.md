# Finite Automata

A finite automata (FA) is an abstract machine that can be used to represent certain forms of computation. Graphically, an FA consists of a number of states and a number of edges between those states.

The machine begin in a start state _S0_. For each input symbol presented to the FA, it moves to the state indicated by the edge with the same label as the input label. Some states of the FA are known as _accepting states_. If the FA is in an _accepting state_ after all input is consumed, then we say that the FA _accepts_ the input. We say that te FA _rejects_ the input string if it ends in a non-accepting state.

Every _Regular Expression_ can be written as a FA, and vice versa.


## Usage

```
r := new(regexp)

// will check the input for 3 occurrencies of the range 'a-c' no matter the order
// will check the input for 1 occurrency of the letter 'o'
// will check the input for 3 occurencies of the letter 'r'
err := r.addState(regexpopts{
    exp:        "a-c",
    recurrency: 3,
}, regexpopts{
    exp:        "o",
    recurrency: 1,
}, regexpopts{
    exp:        "r",
    recurrency: 3,
})

r.evaluate("abcorrr") // true
r.evaluate("bcaorrr") // true
r.evaluate("ior") // false
```