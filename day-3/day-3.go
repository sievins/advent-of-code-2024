package main

import (
	"day-3/data"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	corruptedMemory := data.GetCorruptedMemory()

	total := scan(corruptedMemory)
	fmt.Printf("Total: %d\n", total)

	advancedTotal := advancedScan(corruptedMemory)
	fmt.Printf("Advanced total: %d\n", advancedTotal)
}

func advancedScan(corruptedMemory string) (total int) {
	initialPortionRegex := regexp.MustCompile(`(.*?)don't\(\)(.*)`)
	splitOnFirstDont := initialPortionRegex.FindStringSubmatch(corruptedMemory)

	total += scan(splitOnFirstDont[1])

	remainingEnabledPortionsRegex := regexp.MustCompile(`do\(\).*?(?:don't\(\)|$)`)
	remainingEnabledPortions := remainingEnabledPortionsRegex.FindAllStringSubmatch(splitOnFirstDont[2], -1)

	for _, enabledPortion := range remainingEnabledPortions {
		total += scan(enabledPortion[0])
	}

	return total
}

func scan(corruptedMemory string) (total int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(corruptedMemory, -1)

	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		total += left * right
	}

	return total
}
