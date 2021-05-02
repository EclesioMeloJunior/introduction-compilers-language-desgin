package main

import (
	"errors"
	"strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type exp string

var ErrInvalidExpression = errors.New("invalid expression")

func (e *exp) hasRange() bool {
	return strings.Contains(string(*e), "-")
}

func (e *exp) isValidRange() bool {
	if !e.hasRange() {
		return false
	}

	terms := strings.Split(string(*e), "-")
	if len(terms) != 2 {
		return false
	}

	start, end := terms[0], terms[1]
	if strings.Compare(start, end) != -1 {
		return false
	}

	return true
}

func (e *exp) getRange() string {
	if !e.isValidRange() {
		return ""
	}

	terms := strings.Split(string(*e), "-")
	start, end := strings.Index(alphabet, terms[0]), strings.Index(alphabet, terms[1])

	// if end == z or end == Z then
	if (end+1) >= (len(alphabet)/2) || (end+1) >= len(alphabet) {

		// we get from start til the uppercase Z
		if start >= len(alphabet) {
			return alphabet[start:]
		}

		// otherwise we get til the lowercase z
		return alphabet[start : len(alphabet)/2]
	}

	return alphabet[start : end+1]
}
