package main

import "fmt"

var word = []rune{'X', 'M', 'A', 'S'}

func findNumberOfWords(grid [][]rune) {
	var numberOfWords int

	for y, line := range grid {
		for x := range line {
			numberOfWords += howManyWordsFromCoordinate(x, y, grid)
		}
	}

	fmt.Println("Number of words:", numberOfWords)
}

func howManyWordsFromCoordinate(x int, y int, grid [][]rune) (numberOfWords int) {
	// Assume initial x y coord is valid

	// Word must start with X
	if grid[y][x] != word[0] {
		return 0
	}

	// North
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{0, 0, 0}, []int{-1, -2, -3})
	// North East
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{1, 2, 3}, []int{-1, -2, -3})
	// East
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{1, 2, 3}, []int{0, 0, 0})
	// South East
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{1, 2, 3}, []int{1, 2, 3})
	// South
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{0, 0, 0}, []int{1, 2, 3})
	// South West
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{-1, -2, -3}, []int{1, 2, 3})
	// West
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{-1, -2, -3}, []int{0, 0, 0})
	// North West
	findWordFromCoordinate(&numberOfWords, x, y, grid, []int{-1, -2, -3}, []int{-1, -2, -3})

	return numberOfWords
}

func findWordFromCoordinate(numberOfWords *int, x int, y int, grid [][]rune, xIncrement []int, yIncrement []int) {
	foundWord := true

	for i := 0; i <= 2; i++ {
		currentX := x + xIncrement[i]
		currentY := y + yIncrement[i]

		valid := isValidCoordinate(currentX, currentY, grid)

		if !valid || grid[currentY][currentX] != word[i+1] {
			foundWord = false
			break
		}
	}

	if foundWord {
		*numberOfWords += 1
	}
}

func isValidCoordinate(x int, y int, grid [][]rune) (isValid bool) {
	maxX, maxY := getMaxCoordinates(grid)
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

func getMaxCoordinates(grid [][]rune) (maxX int, maxY int) {
	maxX = len(grid[0]) - 1
	maxY = len(grid) - 1
	return maxX, maxY
}
