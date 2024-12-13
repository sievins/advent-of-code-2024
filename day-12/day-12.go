package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Location struct {
	x int
	y int
}

// Remember which locations have been included in any plot
var completedLocations = make(map[Location]struct{})

// Ignore plants that have already been included in this plot
var currentPlotLocations = make(map[Location]struct{})

func main() {
	// Create 2D slice of runes to represent the garden
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var garden [][]rune
	for _, line := range lines {
		row := []rune(line)
		garden = append(garden, row)
	}

	maxY := len(garden) - 1
	maxX := len(garden[0]) - 1

	price := 0
	discountedPrice := 0

	start := time.Now()

	for y, row := range garden {
		for x := range row {
			location := Location{x, y}
			// Ignore plants that have already been included in any plot
			if _, locationAlreadyIncludedInPlot := completedLocations[location]; locationAlreadyIncludedInPlot {
				continue
			}

			currentPlotLocations = map[Location]struct{}{}
			parameter, area := calculatePlotDimensions(garden, location, maxX, maxY)
			sides := calculateSides(currentPlotLocations)

			price += parameter * area
			discountedPrice += sides * area
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Price:", price)
	fmt.Println("Discounted price:", discountedPrice)
	fmt.Println("Total time:", elapsed)
}

func calculatePlotDimensions(garden [][]rune, location Location, maxX int, maxY int) (parameter int, area int) {
	completedLocations[location] = struct{}{}
	currentPlotLocations[location] = struct{}{}

	nextLocations, locationsParameter := getNextLocations(garden, location, maxX, maxY)

	for _, nextLocation := range nextLocations {
		completedLocations[nextLocation] = struct{}{}
		currentPlotLocations[nextLocation] = struct{}{}
	}

	for _, nextLocation := range nextLocations {
		nextParameter, nextArea := calculatePlotDimensions(garden, nextLocation, maxX, maxY)
		parameter += nextParameter
		area += nextArea
	}

	parameter += locationsParameter
	area += 1

	return parameter, area
}

func getNextLocations(garden [][]rune, location Location, maxX int, maxY int) (nextLocations []Location, parameter int) {
	plant := garden[location.y][location.x]

	possibleLocations := []Location{
		{location.x, location.y - 1}, // up
		{location.x + 1, location.y}, // right
		{location.x, location.y + 1}, // down
		{location.x - 1, location.y}, // left
	}

	for _, possibleLocation := range possibleLocations {
		if _, alreadyAddedPlantToPlot := currentPlotLocations[possibleLocation]; !alreadyAddedPlantToPlot {
			if isInsideGrid(possibleLocation, maxX, maxY) && garden[possibleLocation.y][possibleLocation.x] == plant {
				nextLocations = append(nextLocations, possibleLocation)
			} else {
				parameter++
			}
		}
	}

	return nextLocations, parameter
}

func isInsideGrid(location Location, maxX int, maxY int) bool {
	x := location.x
	y := location.y
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

type Edge struct {
	from, to Location
}

func calculateSides(plotLocations map[Location]struct{}) int {
	edges := transformPlotLocationsIntoEdges(plotLocations)

	// Traverse edges and count corners (same as number of sides)
	// Remove edge after traversing, in order to handle internal edges (circle in a circle)
	var firstEdge Edge
	for key := range edges {
		firstEdge = key
		break
	}
	delete(edges, firstEdge)
	return countCorners(edges, firstEdge, firstEdge, Edge{})
}

func countCorners(edges map[Edge]struct{}, currentEdge Edge, firstEdge Edge, previousEdge Edge) (count int) {
	nextEdge, exists := findNextEdge(edges, currentEdge, previousEdge)

	if !exists {
		// Reached end of edge loop
		if isCorner(currentEdge, firstEdge) {
			// Handle first and last edges in the loop being a corner
			count++
		}
		// There could be inner loops
		if len(edges) > 0 {
			var newFirstEdge Edge
			for key := range edges {
				newFirstEdge = key
				break
			}
			delete(edges, newFirstEdge)
			return count + countCorners(edges, newFirstEdge, newFirstEdge, Edge{})
		}
		return count
	}

	if isCorner(currentEdge, nextEdge) {
		count++
	}

	delete(edges, nextEdge)

	return count + countCorners(edges, nextEdge, firstEdge, currentEdge)
}

func isCorner(edge1 Edge, edge2 Edge) bool {
	edge1Horizontal := edge1.from.y == edge1.to.y
	edge2Horizontal := edge2.from.y == edge2.to.y
	return edge1Horizontal != edge2Horizontal
}

func findNextEdge(edges map[Edge]struct{}, edge Edge, previousEdge Edge) (nextEdge Edge, exists bool) {
	possibleNextEdges := []Edge{}
	for possibleNextEdge := range edges {
		if edgesAreNeighbours(edge, possibleNextEdge) && !edgesAreNeighbours(possibleNextEdge, previousEdge) {
			// If there is a crossroad, never go back
			possibleNextEdges = append(possibleNextEdges, possibleNextEdge)
		}
	}

	if len(possibleNextEdges) == 0 {
		return Edge{}, false
	} else if len(possibleNextEdges) > 1 {
		for _, possibleNextEdge := range possibleNextEdges {
			if isCorner(possibleNextEdge, edge) {
				// If there is a choice - change direction - solves bug which moves across into diagonal section
				return possibleNextEdge, true
			}
		}
	}

	return possibleNextEdges[0], true
}

func edgesAreNeighbours(edge1 Edge, edge2 Edge) bool {
	return edge1.from == edge2.to || edge1.to == edge2.to || edge1.from == edge2.from || edge1.to == edge2.from
}

func transformPlotLocationsIntoEdges(plotLocations map[Location]struct{}) map[Edge]struct{} {
	createEdge := func(l Location, direction string) Edge {
		switch direction {
		case "up":
			return Edge{l, Location{l.x + 1, l.y}}
		case "right":
			return Edge{Location{l.x + 1, l.y}, Location{l.x + 1, l.y + 1}}
		case "down":
			return Edge{Location{l.x, l.y + 1}, Location{l.x + 1, l.y + 1}}
		case "left":
			return Edge{l, Location{l.x, l.y + 1}}
		default:
			return Edge{l, l}
		}
	}

	edges := make(map[Edge]struct{})

	// Edges appear for non adjcent plants
	for plant := range plotLocations {
		x := plant.x
		y := plant.y

		// If no plant above add edge locations
		up := Location{x, y - 1}
		if _, exists := plotLocations[up]; !exists {
			edge := createEdge(plant, "up")
			edges[edge] = struct{}{}
		}

		// If no plant to the right add edge locations
		right := Location{x + 1, y}
		if _, exists := plotLocations[right]; !exists {
			edge := createEdge(plant, "right")
			edges[edge] = struct{}{}
		}

		// If no plant below add edge locations
		down := Location{x, y + 1}
		if _, exists := plotLocations[down]; !exists {
			edge := createEdge(plant, "down")
			edges[edge] = struct{}{}
		}

		// If no plant to the left add edge locations
		left := Location{x - 1, y}
		if _, exists := plotLocations[left]; !exists {
			edge := createEdge(plant, "left")
			edges[edge] = struct{}{}
		}
	}

	return edges
}
