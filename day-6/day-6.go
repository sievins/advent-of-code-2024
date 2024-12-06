package main

import (
	"fmt"
	"slices"
)

var (
	guard       = []rune{'^', '>', 'v', '<'}
	rotateGuard = map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}
)

func main() {
	lab := parse()

	// Part 1
	labCopy := cloneLab(lab)
	getNumberOfPositionsVisited(labCopy)

	// Part 2
	labCopy = cloneLab(lab)
	getNumberOfObstacles(labCopy)
}

func getNumberOfPositionsVisited(lab [][]rune) {
	visited := 1

	x, y := findStartingCoordinate(lab)

	for {
		nextX, nextY := getNextCoordinate(lab, x, y)

		if !isInsideMap(lab, nextX, nextY) {
			break
		}

		if isObstacle(lab, nextX, nextY) {
			lab[y][x] = rotateGuard[lab[y][x]]
			nextX, nextY = getNextCoordinate(lab, x, y)
			if !isInsideMap(lab, nextX, nextY) {
				break
			}
		}

		if lab[nextY][nextX] != 'X' {
			visited += 1
		}

		lab[nextY][nextX] = lab[y][x]
		lab[y][x] = 'X'
		x = nextX
		y = nextY
	}

	fmt.Println("Visited:", visited)
}

func getNumberOfObstacles(lab [][]rune) {
	obstacles := 0
	originalLab := cloneLab(lab)
	x, y := findStartingCoordinate(originalLab)

	for oy, line := range originalLab {
		for ox, item := range line {
			if item == '#' || (ox == x && oy == y) {
				// skip if obstacle already exists or is starting position
				continue
			}

			// add obstacle
			lab[oy][ox] = '#'

			// test
			if gaurdLoops(lab, x, y) {
				obstacles += 1
			}

			// reset lab
			lab = cloneLab(originalLab)
		}
	}

	fmt.Println("Obstacles that cause a loop:", obstacles)
}

func gaurdLoops(lab [][]rune, x int, y int) bool {
	type Coord struct {
		x int
		y int
	}
	positionToDirection := make(map[Coord][]rune)

	for {
		// Does current position include current direction - aka been here before facing the same way
		coord := Coord{x, y}
		if slices.Contains(positionToDirection[coord], lab[y][x]) {
			return true
		}
		// Add current direction to position
		positionToDirection[coord] = append(positionToDirection[coord], lab[y][x])

		nextX, nextY := getNextCoordinate(lab, x, y)
		if !isInsideMap(lab, nextX, nextY) {
			return false
		}

		if isObstacle(lab, nextX, nextY) {
			lab[y][x] = rotateGuard[lab[y][x]]
			nextX, nextY = getNextCoordinate(lab, x, y)

			// Do not walk straight into another obstacle - turn around
			if isObstacle(lab, nextX, nextY) {
				lab[y][x] = rotateGuard[lab[y][x]]
				nextX, nextY = getNextCoordinate(lab, x, y)
			}

			if !isInsideMap(lab, nextX, nextY) {
				return false
			}

			// Add rotated direction to position
			positionToDirection[coord] = append(positionToDirection[coord], lab[y][x])
		}

		lab[nextY][nextX] = lab[y][x]
		x = nextX
		y = nextY
	}
}

func cloneLab(original [][]rune) [][]rune {
	clone := make([][]rune, len(original))

	for i, innerSlice := range original {
		clone[i] = make([]rune, len(innerSlice))
		copy(clone[i], innerSlice)
	}

	return clone
}

func findStartingCoordinate(lab [][]rune) (x int, y int) {
	for yy, line := range lab {
		for xx, item := range line {
			if slices.Contains(guard, item) {
				return xx, yy
			}
		}
	}
	return -1, -1
}

func getNextCoordinate(lab [][]rune, x int, y int) (nextX int, nextY int) {
	directionToNextCoordinate := make(map[rune][]int)
	directionToNextCoordinate['^'] = []int{x, y - 1}
	directionToNextCoordinate['>'] = []int{x + 1, y}
	directionToNextCoordinate['v'] = []int{x, y + 1}
	directionToNextCoordinate['<'] = []int{x - 1, y}
	return directionToNextCoordinate[lab[y][x]][0], directionToNextCoordinate[lab[y][x]][1]
}

func isObstacle(lab [][]rune, x int, y int) bool {
	return lab[y][x] == '#'
}

func isInsideMap(lab [][]rune, x int, y int) (isValid bool) {
	maxX, maxY := getMaxCoordinates(lab)
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

func getMaxCoordinates(lab [][]rune) (maxX int, maxY int) {
	maxX = len(lab[0]) - 1
	maxY = len(lab) - 1
	return maxX, maxY
}
