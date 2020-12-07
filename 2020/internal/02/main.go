package main

import (
	"fmt"

	"github.com/MarkRosemaker/adventofcode/2020/internal/02/data"
)

func main() {
	n := CountValidsOne(data.Input)
	fmt.Printf("Result #1: %d\n", n)

	n = CountValidsTwo(data.Input)
	fmt.Printf("Result #2: %d\n", n)
}

func validateListOne(list []data.PasswordScheme) []bool {
	k := len(list)
	res := make([]bool, k)
	for i := 0; i < k; i++ {
		res[i] = validateOne(list[i])
	}

	return res
}

func validateListTwo(list []data.PasswordScheme) []bool {
	k := len(list)
	res := make([]bool, k)
	for i := 0; i < k; i++ {
		res[i] = validateTwo(list[i])
	}

	return res
}

func validateOne(scheme data.PasswordScheme) bool {
	k := len(scheme.Password)
	count := 0
	for i := 0; i < k; i++ {
		if scheme.Letter == scheme.Password[i] {
			count++
		}
	}
	return count <= scheme.High && count >= scheme.Low
}

func validateTwo(scheme data.PasswordScheme) bool {
	// k := len(scheme.Password)
	count := 0

	// no index out of bounds checks

	if scheme.Password[scheme.Low-1] == scheme.Letter {
		count++
	}
	if scheme.Password[scheme.High-1] == scheme.Letter {
		count++
	}
	return count == 1
}

func CountValidsOne(list []data.PasswordScheme) int {
	return countTrues(validateListOne(list))
}

func CountValidsTwo(list []data.PasswordScheme) int {
	return countTrues(validateListTwo(list))
}

func countTrues(v []bool) int {
	count := 0
	for _, b := range v {
		if b {
			count++
		}
	}
	return count
}
