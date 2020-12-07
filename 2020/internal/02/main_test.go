package main

import (
	"testing"

	"github.com/MarkRosemaker/adventofcode/2020/internal/02/data"
)

func TestValidateList(t *testing.T) {
	res := validateListOne(data.Example)
	if !equal(res, data.ExampleResultOne) {
		t.Fatal("example fails the first test")
	}

	res = validateListTwo(data.Example)
	if !equal(res, data.ExampleResultTwo) {
		t.Fatal("example fails the second test")
	}
}

func TestCoundValids(t *testing.T) {
	n := CountValidsOne(data.Example)
	if n != data.ExampleResultOneCount {
		t.Fatal("example fails the test")
	}

	n = CountValidsTwo(data.Example)
	if n != data.ExampleResultTwoCount {
		t.Fatal("example fails the test")
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
