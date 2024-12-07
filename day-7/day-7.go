package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"
)

func main() {
	equations := parse()

	// Store containing all possible combinations against the number of operators, e.g. 2: [['*', '*'], ['*' '+'], ['+', '*'], ['+', '+']]
	store1 := make(Store)
	store2 := make(Store)

	sumPart1 := 0
	sumPart2 := 0

	// Part 1
	start1 := time.Now()
	for _, equation := range equations {
		operatorCombinations := memoizeCombinations(store1, len(equation.numbers)-1, []Operator{'*', '+'})
		possibleResults := calculate(equation.numbers, operatorCombinations)
		if slices.Contains(possibleResults, equation.result) {
			sumPart1 += equation.result
		}
	}
	elapsed1 := time.Since(start1)

	// Part 2
	start2 := time.Now()
	for _, equation := range equations {
		operatorCombinations := memoizeCombinations(store2, len(equation.numbers)-1, []Operator{'*', '+', ':'})
		possibleResults := calculate(equation.numbers, operatorCombinations)
		if slices.Contains(possibleResults, equation.result) {
			sumPart2 += equation.result
		}
	}
	elapsed2 := time.Since(start2)

	fmt.Println("Part 1:", sumPart1, "Time:", elapsed1)
	fmt.Println("Part 2:", sumPart2, "Time:", elapsed2)
}

func calculate(numbers []int, operatorCombinations [][]Operator) (results []int) {
	for _, operators := range operatorCombinations {
		total := 0
		for i, number := range numbers {
			if i == 0 {
				total = number
			} else {
				operator := operators[i-1]
				switch operator {
				case '+':
					total = total + number
				case '*':
					total = total * number
				case ':':
					totalString := strconv.Itoa(total)
					numberString := strconv.Itoa(number)
					total, _ = strconv.Atoi(totalString + numberString)
				}
			}
		}

		results = append(results, total)
	}

	return results
}
