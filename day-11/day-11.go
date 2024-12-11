package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Parse input
	data, _ := os.ReadFile("data.txt")
	line := strings.TrimSpace(string(data))
	initialStones := strings.Fields(line)

	// The number of stones increases dramatically for each blink.
	// So instead of looping through the new list of stones every blink,
	// store the number of stones there are in a frequency map.
	// e.g. there are 10 stones with the number 3 on it.
	// Then you just have to calculate how a stone with 3 on it changes once,
	// instead of 10 times, and assign all the stones to a new key.
	frequency := make(map[int]int)
	for _, stone := range initialStones {
		num, _ := strconv.Atoi(stone)
		frequency[num]++
	}

	start := time.Now()
	totalStones := blinkblink(frequency, 25)
	elapsed := time.Since(start)
	fmt.Println("Number of stones:", totalStones, "Time:", elapsed)

	start = time.Now()
	totalStones = blinkblink(frequency, 75)
	elapsed = time.Since(start)
	fmt.Println("Number of stones:", totalStones, "Time:", elapsed)
}

func blinkblink(frequency map[int]int, blinks int) (totalStones int) {
	for blink := 0; blink < blinks; blink++ {
		newFrequency := make(map[int]int)

		for stone, count := range frequency {
			if stone == 0 {
				// Turn 0 into 1
				newFrequency[1] += count
			} else if isEvenDigits(stone) {
				// Split the number into two halves
				left, right := splitEven(stone)
				newFrequency[left] += count
				newFrequency[right] += count
			} else {
				// Multiply by 2024
				newFrequency[stone*2024] += count
			}
		}

		frequency = newFrequency
	}

	for _, count := range frequency {
		totalStones += count
	}

	return totalStones
}

// Check if a number has an even number of digits
func isEvenDigits(num int) bool {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}
	return digits%2 == 0
}

// Split a number with an even number of digits into two halves
func splitEven(num int) (int, int) {
	digits := []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}

	mid := len(digits) / 2
	left, right := 0, 0
	for i, d := range digits {
		if i < mid {
			left = left*10 + d
		} else {
			right = right*10 + d
		}
	}

	return left, right
}
