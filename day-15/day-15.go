package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	up Direction = iota
	right
	down
	left
)

type Location struct {
	x, y int
}

type Robot struct {
	direction Direction
	location  Location
}

var runeToDirection = map[rune]Direction{
	'^': up,
	'>': right,
	'v': down,
	'<': left,
}

func printWarehouse(warehouse [][]rune) {
	for _, row := range warehouse {
		fmt.Println(string(row))
	}
}

func printHighlightedRunes(row []rune, highlightIndex int) {
	for i, r := range row {
		if i == highlightIndex {
			// Highlight the rune in red
			fmt.Printf("\033[31m%c\033[0m", r) // Red color escape sequence
		} else {
			// Print the rune normally
			fmt.Printf("%c", r)
		}
	}
	fmt.Println() // Move to the next line
}

func main() {
	data, _ := os.ReadFile("data.txt")
	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	// Part 1

	lines := strings.Split(parts[0], "\n")
	var warehouse [][]rune
	for _, line := range lines {
		warehouse = append(warehouse, []rune(line))
	}

	originalWarehouse := cloneWarehouse(warehouse)

	// printWarehouse(warehouse)

	instructions := []rune(strings.ReplaceAll(parts[1], "\n", ""))

	robot := findRobot(warehouse, instructions)

	for _, instruction := range instructions {
		robot.direction = runeToDirection[instruction]
		movableItems := getMoveableItems(warehouse, robot)
		moveItems(warehouse, &robot, movableItems)
	}

	// printWarehouse(warehouse)

	score := 0
	for y, lines := range warehouse {
		for x, item := range lines {
			if item == 'O' {
				score += 100*y + x
			}
		}
	}

	fmt.Println("GPS score:", score)

	// Part 2

	var wideWarehouse [][]rune
	for _, ys := range originalWarehouse {
		newYs := []rune{}
		for _, item := range ys {
			switch item {
			case '@':
				newYs = append(newYs, '@', '.')
			case 'O':
				newYs = append(newYs, '[', ']')
			default:
				newYs = append(newYs, item, item)
			}
		}
		wideWarehouse = append(wideWarehouse, newYs)
	}

	// printWarehouse(wideWarehouse)

	robot = findRobot(wideWarehouse, instructions)

	for _, instruction := range instructions {
		// printHighlightedRunes(instructions, i)
		robot.direction = runeToDirection[instruction]
		movableItems := getWideMoveableItems(wideWarehouse, robot)
		moveWideItems(wideWarehouse, &robot, movableItems)
		// printWarehouse(wideWarehouse)
	}

	// printWarehouse(wideWarehouse)

	score = 0
	for y, lines := range wideWarehouse {
		for x, item := range lines {
			if item == '[' {
				score += 100*y + x
			}
		}
	}

	fmt.Println("GPS score (part 2):", score)
}

func findRobot(warehouse [][]rune, instructions []rune) (robot Robot) {
	foundRobot := false
	for y, lines := range warehouse {
		if foundRobot {
			break
		}
		for x, item := range lines {
			if item == '@' {
				robot = Robot{
					direction: runeToDirection[instructions[0]],
					location:  Location{x, y},
				}
				foundRobot = true
				break
			}
		}
	}
	return robot
}

func getMoveableItems(warehouse [][]rune, robot Robot) []Location {
	moveable := []Location{
		robot.location,
	}

	// Up
	if robot.direction == up {
		for i := robot.location.y - 1; i >= 0; i-- {
			x := robot.location.x
			y := i
			if warehouse[y][x] == 'O' {
				moveable = append(moveable, Location{x, y})
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = []Location{}
				break
			}
		}
	}

	// Right
	if robot.direction == right {
		for i := robot.location.x + 1; i <= len(warehouse[0])-1; i++ {
			x := i
			y := robot.location.y
			if warehouse[y][x] == 'O' {
				moveable = append(moveable, Location{x, y})
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = []Location{}
				break
			}
		}
	}

	// Down
	if robot.direction == down {
		for i := robot.location.y + 1; i <= len(warehouse)-1; i++ {
			x := robot.location.x
			y := i
			if warehouse[y][x] == 'O' {
				moveable = append(moveable, Location{x, y})
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = []Location{}
				break
			}
		}
	}

	// Left
	if robot.direction == left {
		for i := robot.location.x - 1; i >= 0; i-- {
			x := i
			y := robot.location.y
			if warehouse[y][x] == 'O' {
				moveable = append(moveable, Location{x, y})
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = []Location{}
				break
			}
		}
	}

	return moveable
}

func moveItems(warehouse [][]rune, robot *Robot, movableItems []Location) {
	currentItem := '@'

	for i, moveableItem := range movableItems {
		// Up
		if robot.direction == up {
			// Robot leaves his spot
			if i == 0 {
				warehouse[moveableItem.y][moveableItem.x] = '.'
				robot.location = Location{moveableItem.x, moveableItem.y - 1}
			}
			nextItem := warehouse[moveableItem.y-1][moveableItem.x]
			warehouse[moveableItem.y-1][moveableItem.x] = currentItem
			currentItem = nextItem
		}

		// Right
		if robot.direction == right {
			// Robot leaves his spot
			if i == 0 {
				warehouse[moveableItem.y][moveableItem.x] = '.'
				robot.location = Location{moveableItem.x + 1, moveableItem.y}
			}
			nextItem := warehouse[moveableItem.y][moveableItem.x+1]
			warehouse[moveableItem.y][moveableItem.x+1] = currentItem
			currentItem = nextItem
		}

		// Down
		if robot.direction == down {
			// Robot leaves his spot
			if i == 0 {
				warehouse[moveableItem.y][moveableItem.x] = '.'
				robot.location = Location{moveableItem.x, moveableItem.y + 1}
			}
			nextItem := warehouse[moveableItem.y+1][moveableItem.x]
			warehouse[moveableItem.y+1][moveableItem.x] = currentItem
			currentItem = nextItem
		}

		// Left
		if robot.direction == left {
			// Robot leaves his spot
			if i == 0 {
				warehouse[moveableItem.y][moveableItem.x] = '.'
				robot.location = Location{moveableItem.x - 1, moveableItem.y}
			}
			nextItem := warehouse[moveableItem.y][moveableItem.x-1]
			warehouse[moveableItem.y][moveableItem.x-1] = currentItem
			currentItem = nextItem
		}
	}
}

func getWideMoveableItems(warehouse [][]rune, robot Robot) map[Location]struct{} {
	moveable := map[Location]struct{}{
		robot.location: {},
	}

	// Right and left are simpler - [] can move when ther is no # blocking them
	// Take into account wider walls

	// Right
	if robot.direction == right {
		for i := robot.location.x + 1; i <= len(warehouse[0])-2; i++ {
			x := i
			y := robot.location.y
			if warehouse[y][x] == '[' || warehouse[y][x] == ']' {
				moveable[Location{x, y}] = struct{}{}
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = map[Location]struct{}{}
				break
			}
		}
	}

	// Left
	if robot.direction == left {
		for i := robot.location.x - 1; i >= 1; i-- {
			x := i
			y := robot.location.y
			if warehouse[y][x] == '[' || warehouse[y][x] == ']' {
				moveable[Location{x, y}] = struct{}{}
			}
			if warehouse[y][x] == '.' {
				break
			}
			if warehouse[y][x] == '#' {
				moveable = map[Location]struct{}{}
				break
			}
		}
	}

	// Up
	if robot.direction == up {
		// Set of columns pushing up from
		xs := make(map[int]struct{})
		nextxs := make(map[int]struct{})
		nextxs[robot.location.x] = struct{}{}

		for i := robot.location.y - 1; i >= 0; i-- {
			y := i

			xs = nextxs
			nextxs = make(map[int]struct{})

			allSpaces := true
			anyWall := false

			for x := range xs {
				if warehouse[y][x] == '[' {
					moveable[Location{x, y}] = struct{}{}
					moveable[Location{x + 1, y}] = struct{}{}
					nextxs[x] = struct{}{}
					nextxs[x+1] = struct{}{}
				}
				if warehouse[y][x] == ']' {
					moveable[Location{x, y}] = struct{}{}
					moveable[Location{x - 1, y}] = struct{}{}
					nextxs[x] = struct{}{}
					nextxs[x-1] = struct{}{}
				}
				if warehouse[y][x] != '.' {
					allSpaces = false
				}
				if warehouse[y][x] == '#' {
					anyWall = true
					break
				}
			}

			if allSpaces {
				break
			}

			if anyWall {
				moveable = map[Location]struct{}{}
				break
			}
		}
	}

	// Down
	if robot.direction == down {
		// Set of columns pushing down from
		xs := make(map[int]struct{})
		nextxs := make(map[int]struct{})
		nextxs[robot.location.x] = struct{}{}

		for i := robot.location.y + 1; i <= len(warehouse)-1; i++ {
			y := i

			xs = nextxs
			nextxs = make(map[int]struct{})

			allSpaces := true
			anyWall := false

			for x := range xs {
				if warehouse[y][x] == '[' {
					moveable[Location{x, y}] = struct{}{}
					moveable[Location{x + 1, y}] = struct{}{}
					nextxs[x] = struct{}{}
					nextxs[x+1] = struct{}{}
				}
				if warehouse[y][x] == ']' {
					moveable[Location{x, y}] = struct{}{}
					moveable[Location{x - 1, y}] = struct{}{}
					nextxs[x] = struct{}{}
					nextxs[x-1] = struct{}{}
				}
				if warehouse[y][x] != '.' {
					allSpaces = false
				}
				if warehouse[y][x] == '#' {
					anyWall = true
					break
				}
			}

			if allSpaces {
				break
			}

			if anyWall {
				moveable = map[Location]struct{}{}
				break
			}
		}
	}

	return moveable
}

func moveWideItems(warehouse [][]rune, robot *Robot, movableItems map[Location]struct{}) {
	// Moveable items are not orders so have to handle moving in any order.
	// Put item in new position. If item behind is in list of moveable items, put in old spot, otherwise put '.' in old spot.
	// Explicitally update robots location.

	unmovedWarehouse := cloneWarehouse(warehouse)

	for item := range movableItems {
		x := item.x
		y := item.y

		// Up
		if robot.direction == up {
			if unmovedWarehouse[y][x] == '@' {
				robot.location = Location{x, y - 1}
			}
			warehouse[y-1][x] = unmovedWarehouse[y][x]
			if _, isItemBehind := movableItems[Location{x, y + 1}]; isItemBehind {
				warehouse[y][x] = unmovedWarehouse[y+1][x]
			} else {
				warehouse[y][x] = '.'
			}
		}

		// Down
		if robot.direction == down {
			if unmovedWarehouse[y][x] == '@' {
				robot.location = Location{x, y + 1}
			}
			warehouse[y+1][x] = unmovedWarehouse[y][x]
			if _, isItemBehind := movableItems[Location{x, y - 1}]; isItemBehind {
				warehouse[y][x] = unmovedWarehouse[y-1][x]
			} else {
				warehouse[y][x] = '.'
			}
		}

		// Right
		if robot.direction == right {
			if unmovedWarehouse[y][x] == '@' {
				robot.location = Location{x + 1, y}
			}
			warehouse[y][x+1] = unmovedWarehouse[y][x]
			if _, isItemBehind := movableItems[Location{x - 1, y}]; isItemBehind {
				warehouse[y][x] = unmovedWarehouse[y][x-1]
			} else {
				warehouse[y][x] = '.'
			}
		}

		// Left
		if robot.direction == left {
			if unmovedWarehouse[y][x] == '@' {
				robot.location = Location{x - 1, y}
			}
			warehouse[y][x-1] = unmovedWarehouse[y][x]
			if _, isItemBehind := movableItems[Location{x + 1, y}]; isItemBehind {
				warehouse[y][x] = unmovedWarehouse[y][x+1]
			} else {
				warehouse[y][x] = '.'
			}
		}
	}
}

func cloneWarehouse(original [][]rune) [][]rune {
	clone := make([][]rune, len(original))
	for i := range original {
		clone[i] = make([]rune, len(original[i]))
		copy(clone[i], original[i])
	}
	return clone
}
