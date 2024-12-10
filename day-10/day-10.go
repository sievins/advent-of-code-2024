package main

import (
	"fmt"
	"time"
)

type Location struct {
	x int
	y int
}

// Remember locations so that we can break early should another trail already have traversed here
var startToVisitedLocations = make(map[Location]map[Location]struct{})

func main() {
	grid := parse()

	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	// Part 1

	start := time.Now()

	sumits := 0
	for y, xs := range grid {
		for x, height := range xs {
			if height == 0 {
				sumits += findSumits(Location{x: x, y: y}, grid, maxX, maxY, height, Location{x: x, y: y})
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Sumits:", sumits, "Time:", elapsed)

	// Part 2

	start = time.Now()

	score := 0
	for y, xs := range grid {
		for x, height := range xs {
			if height == 0 {
				score += findScore(Location{x: x, y: y}, grid, maxX, maxY, height, Location{x: x, y: y})
			}
		}
	}

	elapsed = time.Since(start)

	fmt.Println("Score:", score, "Time:", elapsed)
}

func findSumits(location Location, grid [][]int, maxX int, maxY int, height int, start Location) (count int) {
	if _, visited := startToVisitedLocations[start][location]; !visited && height == 9 {
		markVisited(start, location)
		return 1
	}

	markVisited(start, location)

	nextLocations := getNextLocations(location, grid, maxX, maxY, height, start, true)
	for _, nextLocation := range nextLocations {
		count += findSumits(nextLocation, grid, maxX, maxY, height+1, start)
	}

	return count
}

func findScore(location Location, grid [][]int, maxX int, maxY int, height int, start Location) (count int) {
	if height == 9 {
		return 1
	}

	nextLocations := getNextLocations(location, grid, maxX, maxY, height, start, false)
	for _, nextLocation := range nextLocations {
		count += findScore(nextLocation, grid, maxX, maxY, height+1, start)
	}

	return count
}

// Returns adjacent locations that are valid locations, haven't been visited yet and are one up
func getNextLocations(location Location, grid [][]int, maxX int, maxY int, height int, start Location, hasToBeUnique bool) (locations []Location) {
	possibleLocations := []Location{
		{x: location.x, y: location.y - 1}, // up
		{x: location.x + 1, y: location.y}, // right
		{x: location.x, y: location.y + 1}, // down
		{x: location.x - 1, y: location.y}, // left
	}

	for _, possibleLocation := range possibleLocations {
		if isInsideGrid(possibleLocation, maxX, maxY) {
			if grid[possibleLocation.y][possibleLocation.x] == height+1 {
				if !hasToBeUnique {
					locations = append(locations, possibleLocation)
				} else if _, visited := startToVisitedLocations[start][possibleLocation]; !visited {
					locations = append(locations, possibleLocation)
				}
			}
		}
	}

	return locations
}

func isInsideGrid(location Location, maxX int, maxY int) bool {
	x := location.x
	y := location.y
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

func markVisited(start, location Location) {
	// Ensure the nested map is initialized
	if startToVisitedLocations[start] == nil {
		startToVisitedLocations[start] = make(map[Location]struct{})
	}

	// Mark the location as visited
	startToVisitedLocations[start][location] = struct{}{}
}
