package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	numbers []int
	result  int
}

func parse() (equations []Equation) {
	data, _ := os.ReadFile("data.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])
		numbersAsStrings := strings.Split(parts[1], " ")
		numbers, _ := convertStringsToInts(numbersAsStrings)
		equation := Equation{
			result:  result,
			numbers: numbers,
		}
		equations = append(equations, equation)
	}

	return equations
}

func convertStringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q at index %d: %v", s, i, err)
		}
		ints[i] = n
	}
	return ints, nil
}
