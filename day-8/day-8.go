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

func main() {
	data, _ := os.ReadFile("data.txt")
	ys := strings.Split(strings.TrimSpace(string(data)), "\n")

	start := time.Now()

	maxX := len(ys[0]) - 1
	maxY := len(ys) - 1

	isInsideMap := func(location Location) bool {
		x := location.x
		y := location.y
		return x >= 0 && x <= maxX && y >= 0 && y <= maxY
	}

	antennas := make(map[rune][]Location)
	// Unique set of antinode locations
	antinodes := make(map[Location]struct{})
	// Another unique set of antinodes for part 2
	antinodes2 := make(map[Location]struct{})

	for y, xs := range ys {
		for x, letter := range xs {
			if letter == '.' {
				continue
			}

			if antennas[letter] == nil {
				// Add first letter to antennas list
				antennas[letter] = append(antennas[letter], Location{x, y})
			} else {
				// When there are more than 1 antennas find the antinode locations
				for _, antennaLocation := range antennas[letter] {
					dx := antennaLocation.x - x
					dy := antennaLocation.y - y

					// Part 1: Find to antinodes either sides of the antennas
					firstAntinodeLocation := Location{
						x: x - dx,
						y: y - dy,
					}
					secondAntinodeLocation := Location{
						x: x + 2*dx,
						y: y + 2*dy,
					}
					if isInsideMap(firstAntinodeLocation) {
						antinodes[firstAntinodeLocation] = struct{}{}
					}
					if isInsideMap(secondAntinodeLocation) {
						antinodes[secondAntinodeLocation] = struct{}{}
					}

					// Part 2: Find multiple antinodes, at same frequency, along the line of the antennas
					// Add starting antennas
					antinodes2[Location{x: x, y: y}] = struct{}{}
					px := x + dx
					py := y + dy
					nx := x - dx
					ny := y - dy
					for isInsideMap(Location{x: px, y: py}) {
						antinodes2[Location{x: px, y: py}] = struct{}{}
						px = px + dx
						py = py + dy
					}
					for isInsideMap(Location{x: nx, y: ny}) {
						antinodes2[Location{x: nx, y: ny}] = struct{}{}
						nx = nx - dx
						ny = ny - dy
					}
				}

				// Add the letter to antennas list
				antennas[letter] = append(antennas[letter], Location{x, y})
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Number of antinodes - part 1:", len(antinodes))
	fmt.Println("Number of antinodes - part 2:", len(antinodes2))
	fmt.Println("Total time elapsed:", elapsed)
}
