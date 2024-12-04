package main

import "fmt"

func findNumberOfMASPatterns(grid [][]rune) {
	var numberOfPatterns int

	for y, line := range grid {
		for x := range line {
			if isPatternAtCoordinate(x, y, grid) {
				numberOfPatterns += 1
			}
		}
	}

	fmt.Println("Number of X-MAS patterns:", numberOfPatterns)
}

func isPatternAtCoordinate(x int, y int, grid [][]rune) (found bool) {
	// Assume initial x y coord is valid

	// Pattern must start with A
	if grid[y][x] != 'A' {
		return false
	}

	if canCreateXShapeFromCoordinate(x, y, grid) {
		ne := grid[y-1][x+1]
		se := grid[y+1][x+1]
		sw := grid[y+1][x-1]
		nw := grid[y-1][x-1]

		if ((nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')) && ((ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M')) {
			return true
		}
	}

	return false
}

func canCreateXShapeFromCoordinate(x int, y int, grid [][]rune) (isValid bool) {
	neX := x + 1
	neY := y - 1
	seX := x + 1
	seY := y + 1
	swX := x - 1
	swY := y + 1
	nwX := x - 1
	nwY := y - 1

	return isValidCoordinate(neX, neY, grid) &&
		isValidCoordinate(seX, seY, grid) &&
		isValidCoordinate(swX, swY, grid) &&
		isValidCoordinate(nwX, nwY, grid)
}
