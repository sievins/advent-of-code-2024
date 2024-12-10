package main

import (
	"os"
	"strings"
)

func parse() (grid [][]int) {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		row := stringToIntSlice(line)
		grid = append(grid, row)
	}

	return grid
}

func stringToIntSlice(input string) []int {
	result := make([]int, len(input))
	for i, char := range input {
		result[i] = int(char - '0')
	}
	return result
}
