package main

import (
	"errors"
	"fmt"

	"github.com/MarkRosemaker/adventofcode/2020/internal/01/data"
)

func main() {
	res, err := ProblemOne(data.Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for Problem #1: %d\n", res)

	res, err = ProblemTwo(data.Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for Problem #2: %d\n", res)
}

func ProblemOne(list []int) (int, error) {
	first, second, err := findPair(list)
	if err != nil {
		return 0, err
	}
	return first * second, nil
}

func ProblemTwo(list []int) (int, error) {
	first, second, third, err := findTriplet(list)
	if err != nil {
		return 0, err
	}
	return first * second * third, nil
}

func findPair(list []int) (int, int, error) {
	k := len(list)
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if list[i]+list[j] == data.Sum {
				return list[i], list[j], nil
			}
		}
	}
	return 0, 0, errors.New("not found")
}

func findTriplet(list []int) (int, int, int, error) {
	k := len(list)
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			for l := j + 1; l < k; l++ {
				if list[i]+list[j]+list[l] == data.Sum {
					return list[i], list[j], list[l], nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("not found")
}
