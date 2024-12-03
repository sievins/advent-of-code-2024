package main

import (
	"day-1/data"
	"fmt"
	"sort"
)

func main() {
	cols := data.Data()
	col1, col2 := data.Convert(cols)

	summedDifference := sumDifference(col1, col2)
	fmt.Printf("The summed difference is %d\n", summedDifference)

	similarityScore := similarity(col1, col2)
	fmt.Printf("The similarity score is %d\n", similarityScore)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sumDifference(leftList []int, rightList []int) (sum int) {
	sort.Ints(leftList)
	sort.Ints(rightList)

	for index, item := range leftList {
		diff := absInt(item - rightList[index])
		sum += diff
	}

	return sum
}

func similarity(leftList []int, rightList []int) (score int) {
	frequencyOfNumber := make(map[int]int)

	for _, item := range rightList {
		if count, exists := frequencyOfNumber[item]; exists {
			frequencyOfNumber[item] = count + 1
		} else {
			frequencyOfNumber[item] = 1
		}
	}

	for _, item := range leftList {
		if count, exists := frequencyOfNumber[item]; exists {
			score += count * item
		}
	}

	return score
}
