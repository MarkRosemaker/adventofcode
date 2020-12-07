package main

import (
	"testing"

	"github.com/MarkRosemaker/adventofcode/2020/internal/01/data"
)

func TestFindPair(t *testing.T) {

	first, second, err := findPair(data.Example)
	if err != nil {
		t.Fatal(err)
	}

	if sum := first + second; sum != data.Sum {
		t.Fatalf("sum doesn't match: %d + %d = %d", first, second, sum)
	}

	first, second, err = findPair(data.Input)
	if err != nil {
		t.Fatal(err)
	}

	if sum := first + second; sum != data.Sum {
		t.Fatalf("sum doesn't match: %d + %d = %d", first, second, sum)
	}

}

func TestFindTriplet(t *testing.T) {

	first, second, third, err := findTriplet(data.Example)
	if err != nil {
		t.Fatal(err)
	}

	if sum := first + second + third; sum != data.Sum {
		t.Fatalf("sum doesn't match: %d + %d = %d", first, second, sum)
	}

	first, second, third, err = findTriplet(data.Input)
	if err != nil {
		t.Fatal(err)
	}

	if sum := first + second + third; sum != data.Sum {
		t.Fatalf("sum doesn't match: %d + %d = %d", first, second, sum)
	}

}

func TestProblemOne(t *testing.T) {

	res, err := ProblemOne(data.Example)
	if err != nil {
		t.Fatal(err)
	}

	if res != data.ExampleResultOne {
		t.Errorf("result doesn't match the one in the example: %d", res)
	}
}

func TestProblemTwo(t *testing.T) {

	res, err := ProblemTwo(data.Example)
	if err != nil {
		t.Fatal(err)
	}

	if res != data.ExampleResultTwo {
		t.Errorf("result doesn't match the one in the example: %d", res)
	}
}
