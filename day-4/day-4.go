package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := create2DGrid()

	findNumberOfWords(grid)
	findNumberOfMASPatterns(grid)
}

func create2DGrid() (grid [][]rune) {
	data, err := os.ReadFile("word-search.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		row := []rune(line)
		grid = append(grid, row)
	}

	return grid
}
