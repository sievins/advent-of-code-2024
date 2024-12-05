package main

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	rules, updates := parse()

	validUpdates, invalidUpdates := getValidAndInvalidUpdates(rules, updates)

	sumOfValidMiddlePageNumbers := sumMiddlePageNumbersOfValidUpdates(validUpdates)
	fmt.Println("Valid total:", sumOfValidMiddlePageNumbers)

	sumOfInvalidMiddlePageNumbers := sumMiddlePageNumbersOfInvalidUpdates(invalidUpdates, rules)
	fmt.Println("Invalid total:", sumOfInvalidMiddlePageNumbers)
}

func getValidAndInvalidUpdates(rules []string, updates [][]string) (validUpdates [][]string, invalidUpdates [][]string) {
	for _, update := range updates {
		if testIsUpdateValid(update, rules) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	return validUpdates, invalidUpdates
}

func testIsUpdateValid(update []string, rules []string) bool {
	isUpdateValid := true

	for i, page := range update {
		if i == len(update)-1 {
			continue
		}

		subsequentPages := update[i+1:]
		for _, subsubsequentPage := range subsequentPages {
			if !slices.Contains(rules, page+"|"+subsubsequentPage) {
				isUpdateValid = false
				break
			}
		}

		if !isUpdateValid {
			break
		}
	}

	return isUpdateValid
}

func sumMiddlePageNumbersOfValidUpdates(validUpdates [][]string) int {
	sumOfMiddlePageNumbers := 0
	for _, validUpdate := range validUpdates {
		middlePage := validUpdate[(len(validUpdate)-1)/2]
		middleNumber, _ := strconv.Atoi(middlePage)
		sumOfMiddlePageNumbers += middleNumber
	}
	return sumOfMiddlePageNumbers
}

func sumMiddlePageNumbersOfInvalidUpdates(invalidUpdates [][]string, rules []string) (sumOfMiddlePageNumbers int) {
	for _, invalidUpdate := range invalidUpdates {
		sumOfMiddlePageNumbers += getMiddleNumberFromInvalidUpdate(invalidUpdate, rules)
	}
	return sumOfMiddlePageNumbers
}

func getMiddleNumberFromInvalidUpdate(update []string, rules []string) (middleNumber int) {
	positionToPage := make(map[int]string)

	for i, page := range update {
		otherPages := make([]string, 0, len(update)-1)
		otherPages = append(otherPages, update[:i]...)
		otherPages = append(otherPages, update[i+1:]...)

		position := 0
		for _, otherPage := range otherPages {
			if slices.Contains(rules, otherPage+"|"+page) {
				position += 1
			}
		}
		positionToPage[position] = page
	}

	middlePage := positionToPage[(len(update)-1)/2]
	middleNumber, _ = strconv.Atoi(middlePage)
	return middleNumber
}
