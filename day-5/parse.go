package main

import (
	"os"
	"strings"
)

func parse() (rules []string, updates [][]string) {
	data, _ := os.ReadFile("data.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		if isRule(line) {
			rules = append(rules, line)
		} else if isPageNumbers(line) {
			update := strings.Split(line, ",")
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func isRule(s string) bool {
	return strings.Contains(s, "|")
}

func isPageNumbers(s string) bool {
	return !isRule(s) && s != ""
}
