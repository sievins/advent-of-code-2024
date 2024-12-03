package main

import (
	"day-2/data"
	"fmt"
)

func main() {
	input := data.Data()
	reports := data.Convert(input)

	numberOfSafeReports := calculateNumberOfSafeReports(reports)
	fmt.Printf("The number of safe reports are %d\n", numberOfSafeReports)

	numberOfSafeIshReports := calculateNumberOfSafeIshReports(reports)
	fmt.Printf("The number of safe ish reports are %d\n", numberOfSafeIshReports)
}

func calculateNumberOfSafeIshReports(reports [][]int) (numberOfSafeIshReports int) {
	for _, report := range reports {
		if isReportSafe(report) {
			numberOfSafeIshReports += 1
		} else {
			var isReportWithLevelRemovedSafe bool
			for i := range report {
				clonedReport := append([]int(nil), report...)
				reportWithLevelRemoved := append(clonedReport[:i], clonedReport[i+1:]...)
				if isReportSafe(reportWithLevelRemoved) {
					isReportWithLevelRemovedSafe = true
				}
			}
			if isReportWithLevelRemovedSafe {
				numberOfSafeIshReports += 1
			}
		}
	}
	return numberOfSafeIshReports
}

func calculateNumberOfSafeReports(reports [][]int) (numberOfSafeReports int) {
	for _, report := range reports {
		if isReportSafe(report) {
			numberOfSafeReports += 1
		}
	}

	return numberOfSafeReports
}

func isReportSafe(report []int) (isSafe bool) {
	isIncreasing := report[0] < report[1]
	isDecreasing := report[0] > report[1]

	if !isIncreasing && !isDecreasing {
		return false
	}

	if isIncreasing {
		for i, field := range report[:len(report)-1] {
			if field >= report[i+1] || field+3 < report[i+1] {
				return false
			}
		}
	} else {
		for i, field := range report[:len(report)-1] {
			if field <= report[i+1] || field-3 > report[i+1] {
				return false
			}
		}
	}

	return true
}
