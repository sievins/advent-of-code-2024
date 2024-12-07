package main

type Operator rune

const (
	Times Operator = '*'
	Plus  Operator = '+'
	Sum   Operator = ':'
)

type Store map[int][][]Operator

func memoizeCombinations(store Store, count int, operators []Operator) [][]Operator {
	if combinations, exists := store[count]; exists {
		return combinations
	}

	newCombinations := generateCombinations(count, operators)
	store[count] = newCombinations
	return newCombinations
}

func generateCombinations(count int, operators []Operator) [][]Operator {
	var results [][]Operator

	// Recursive function to build combinations
	var generate func(current []Operator, depth int)
	generate = func(current []Operator, depth int) {
		if depth == count {
			// Add a copy of the current slice to results
			combination := make([]Operator, len(current))
			copy(combination, current)
			results = append(results, combination)
			return
		}

		// Add '*' and '+' and '∑' and recurse
		for _, operator := range operators {
			generate(append(current, operator), depth+1)
		}
	}

	// Start recursion
	generate([]Operator{}, 0)

	return results
}
