package finiteautomata

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidRecurrency = errors.New("invalid recurrency, must be equal or greater then 0")

type regexp struct {
	counter int
	state   *regexpstate
	last    *regexpstate
}

type regexpstate struct {
	opts regexpopts
	next *regexpstate
}

type regexpopts struct {
	exp        exp
	recurrency int
}

// add a new state the the chain of
func (r *regexp) addState(opts ...regexpopts) error {
	for _, opt := range opts {
		if opt.exp.hasRange() && !opt.exp.isValidRange() {
			return ErrInvalidExpression
		}

		if opt.recurrency < 0 {
			return ErrInvalidRecurrency
		}

		s := &regexpstate{
			opts: opt,
			next: nil,
		}

		if r.state == nil {
			r.state, r.last = s, s
		} else {
			r.last.next = s
			r.last = s
		}

		r.counter++
	}

	return nil
}

func (r *regexp) evaluate(input string) bool {
	if r.state == nil {
		return false
	}

	curr := r.state
	recurrency := 0

	for _, s := range input {
		if recurrency == 0 {
			recurrency = curr.opts.recurrency
		}

		if curr.opts.exp.hasRange() {
			exprange := curr.opts.exp.getRange()

			// check if expression range contains the input
			if !strings.Contains(exprange, string(s)) {
				return false
			}
		} else {
			if strings.Compare(string(s), string(curr.opts.exp)) != 0 {
				return false
			}
		}

		recurrency--
		if recurrency != 0 {
			continue
		}

		if curr.next == nil {
			break
		}

		curr = curr.next
	}

	return curr == r.last
}

// print my state
func (r *regexp) print() {
	if r.state == nil {
		fmt.Println("set of states are empty")
		return
	}

	s := r.state
	for {
		fmt.Printf("%s (%v) -> ", s.opts.exp, s.opts.recurrency)

		if s.next == nil {
			fmt.Println(" end")
			break
		}

		s = s.next
	}
}
