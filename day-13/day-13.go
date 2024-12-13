package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
	x1, x2, y1, y2, z1, z2 int
}

func (e Equation) Solve() (tokens int) {
	x1 := e.x1
	x2 := e.x2
	y1 := e.y1
	y2 := e.y2
	z1 := e.z1
	z2 := e.z2

	determinant := x1*y2 - y1*x2

	// When the determinant is 0, either the lines are the same (dependant) or parallel (inconsistent)
	if determinant == 0 {
		panic("It seems that this edge case isn't in my data set")
	}

	coefficient := y2*z1 - x2*z2

	// Check if A is an integer
	if coefficient%determinant != 0 {
		return 0
	}

	// Compute A
	a := coefficient / determinant

	// Check if B is an integer
	numeratorB := z1 - x1*a
	if numeratorB < 0 || numeratorB%x2 != 0 {
		return 0
	}

	// Compute B
	b := numeratorB / x2

	// Ensure both A and B are positive
	if a < 0 || b < 0 {
		return 0
	}

	return 3*a + b
}

var (
	buttonRegex = regexp.MustCompile(`Button [A|B]: X\+(\d+), Y\+(\d+)`)
	prizeRegex  = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

func main() {
	data, _ := os.ReadFile("data.txt")
	instructions := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	start := time.Now()

	total1 := 0
	total2 := 0

	for _, instruction := range instructions {
		equation1 := Equation{}
		equation2 := Equation{}

		lines := strings.Split(instruction, "\n")
		for i, line := range lines {
			// Button A
			if i == 0 {
				matches := buttonRegex.FindStringSubmatch(line)
				x1, _ := strconv.Atoi(matches[1])
				y1, _ := strconv.Atoi(matches[2])
				equation1.x1 = x1
				equation1.y1 = y1
				equation2.x1 = x1
				equation2.y1 = y1
			}

			// Button B
			if i == 1 {
				matches := buttonRegex.FindStringSubmatch(line)
				x2, _ := strconv.Atoi(matches[1])
				y2, _ := strconv.Atoi(matches[2])
				equation1.x2 = x2
				equation1.y2 = y2
				equation2.x2 = x2
				equation2.y2 = y2
			}

			// Prize
			if i == 2 {
				matches := prizeRegex.FindStringSubmatch(line)
				z1, _ := strconv.Atoi(matches[1])
				z2, _ := strconv.Atoi(matches[2])
				equation1.z1 = z1
				equation1.z2 = z2
				equation2.z1 = z1 + 10000000000000
				equation2.z2 = z2 + 10000000000000
			}
		}

		tokens1 := equation1.Solve()
		tokens2 := equation2.Solve()

		total1 += tokens1
		total2 += tokens2
	}

	elapsed := time.Since(start)

	fmt.Println("Total (part 1):", total1)
	fmt.Println("Total (part 2):", total2)
	fmt.Println("Total time:", elapsed)
}
