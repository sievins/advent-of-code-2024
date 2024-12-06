package main

import (
	"os"
	"strings"
)

func parse() (lab [][]rune) {
	data, _ := os.ReadFile("data.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		row := []rune(line)
		lab = append(lab, row)
	}

	return lab
}
