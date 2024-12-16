package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

// Maze representation
type Maze [][]rune

func (m Maze) at(l Location) rune {
	return m[l.y][l.x]
}

// Location represents a position in the maze
type Location struct {
	x, y int
}

// Direction enums
type Direction int

const (
	north Direction = iota
	east
	south
	west
)

// State represents a position, direction, and score
type State struct {
	location  Location
	direction Direction
	score     int
}

// PriorityQueue implements a min-heap for BFS paths
type PriorityQueue []State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// BFS solution
func bfs(maze Maze, start State) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, start)

	// Set of visited states
	visited := make(map[string]bool)

	// Track the lowest score of completed routes
	lowestScore := int(^uint(0) >> 1) // Max int

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		// If we've already found a better path, skip this one
		if current.score >= lowestScore {
			continue
		}

		// If we reached 'E', update the lowest score
		if maze.at(current.location) == 'E' {
			if current.score < lowestScore {
				lowestScore = current.score
			}
			continue
		}

		// Explore all possible moves
		for _, dir := range []Direction{north, east, south, west} {
			nextLoc := moveInDirection(current.location, dir)
			if maze.at(nextLoc) == '#' {
				continue
			}

			// Create a unique key for the visited state
			stateKey := fmt.Sprintf("%d,%d,%d", nextLoc.x, nextLoc.y, dir)
			if visited[stateKey] {
				continue
			}

			// Calculate new score
			newScore := current.score + 1 // Moving forward
			if current.direction != dir {
				newScore += 1000 // Turning 90 degrees
			}

			nextState := State{
				location:  nextLoc,
				direction: dir,
				score:     newScore,
			}

			// Mark this state as visited and push it to the queue
			visited[stateKey] = true
			heap.Push(pq, nextState)
		}
	}

	return lowestScore
}

func moveInDirection(loc Location, dir Direction) Location {
	switch dir {
	case north:
		return Location{loc.x, loc.y - 1}
	case east:
		return Location{loc.x + 1, loc.y}
	case south:
		return Location{loc.x, loc.y + 1}
	case west:
		return Location{loc.x - 1, loc.y}
	default:
		return loc
	}
}

func getStart(maze Maze) State {
	for y, line := range maze {
		for x, item := range line {
			if item == 'S' {
				return State{
					location:  Location{x, y},
					direction: east, // Starts facing east
					score:     0,    // Initial score is 0
				}
			}
		}
	}
	return State{}
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	maze := Maze{}
	for _, line := range lines {
		maze = append(maze, []rune(line))
	}

	// Find start and solve the maze
	start := getStart(maze)
	lowestScore := bfs(maze, start)

	fmt.Println("Lowest score:", lowestScore)
}
