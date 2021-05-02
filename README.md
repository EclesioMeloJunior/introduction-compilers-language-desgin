# Finite Automata

A finite automata (FA) is an abstract machine that can be used to represent certain forms of computation. Graphically, an FA consists of a number of states and a number of edges between those states.

The machine begin in a start state _S0_. For each input symbol presented to the FA, it moves to the state indicated by the edge with the same label as the input label. Some states of the FA are known as _accepting states_. If the FA is in an _accepting state_ after all input is consumed, then we say that the FA _accepts_ the input. We say that te FA _rejects_ the input string if it ends in a non-accepting state.

Every _Regular Expression_ can be written as a FA, and vice versa.