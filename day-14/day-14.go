package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

type Vector struct {
	x, y int
}

type SafestSecond struct {
	seconds, safetyFactor int
}

var instructionRegex = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

const (
	width  = 101
	height = 103
)

const (
	middleColumn = (width - 1) / 2
	middleRow    = (height - 1) / 2
)

func main() {
	data, _ := os.ReadFile("data.txt")
	instructions := strings.Split(strings.TrimSpace(string(data)), "\n")

	finalPositions := []Position{}
	secondsToFinalPositions := make(map[int][]Position)

	for _, instruction := range instructions {
		// Part 1

		matches := instructionRegex.FindStringSubmatch(instruction)
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])

		finalPosition := calculateFinalPosition(Position{px, py}, Vector{vx, vy}, 100)
		finalPositions = append(finalPositions, finalPosition)

		// Part 2

		for i := 1; i < 10000; i++ {
			finalPosition := calculateFinalPosition(Position{px, py}, Vector{vx, vy}, i)
			secondsToFinalPositions[i] = append(secondsToFinalPositions[i], finalPosition)
		}
	}

	// Part 1

	safetyFactor := calculateSafetyFactor(finalPositions)
	fmt.Println("Safety factor:", safetyFactor)

	// Part 2

	safestSecond := SafestSecond{
		seconds:      0,
		safetyFactor: math.MaxInt,
	}

	for seconds, positions := range secondsToFinalPositions {
		safetyFactor := calculateSafetyFactor(positions)

		if safetyFactor < safestSecond.safetyFactor {
			safestSecond = SafestSecond{
				seconds:      seconds,
				safetyFactor: safetyFactor,
			}
		}
	}

	fmt.Println("Safest second:", safestSecond.seconds)
}

func calculateFinalPosition(p Position, v Vector, s int) Position {
	x := (p.x + s*v.x) % width
	if x < 0 {
		x = width + x
	}

	y := (p.y + s*v.y) % height
	if y < 0 {
		y = height + y
	}

	return Position{x, y}
}

func calculateSafetyFactor(positions []Position) (saftetyFactor int) {
	topLeftCount := 0
	topRightCount := 0
	bottomLeftCount := 0
	bottomRightCount := 0

	for _, position := range positions {
		if position.x < middleColumn {
			// Left
			if position.y < middleRow {
				// Top
				topLeftCount++
			} else if position.y > middleRow {
				// Bottom
				bottomLeftCount++
			}
		} else if position.x > middleColumn {
			// Right
			if position.y < middleRow {
				// Top
				topRightCount++
			} else if position.y > middleRow {
				// Bottom
				bottomRightCount++
			}
		}
	}

	return topLeftCount * topRightCount * bottomLeftCount * bottomRightCount
}
