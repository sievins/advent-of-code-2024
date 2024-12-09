package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")
	line := []rune(lines[0])

	// Part 1

	start := time.Now()

	var blocks []int
	isFile := true
	fileID := 0

	for len(line) > 0 {
		if isFile {
			blocks = append(blocks, fillSlice(fileID, toInt(line[0]))...)
			line = line[1:]

			isFile = !isFile
		} else {
			freeCount := toInt(line[0])
			lastFileCount := toInt(line[len(line)-1])
			lastFileID := len(line)/2 + fileID

			if freeCount < lastFileCount {
				// blocks: Fill all the free space with the last fileID. line: Remove the free space from beginning, leave the remaining filesIDs at the end, iterate.
				blocks = append(blocks, fillSlice(lastFileID, freeCount)...)
				line[len(line)-1] = toRune(lastFileCount - freeCount)
				line = line[1:]

				isFile = !isFile
				fileID++
			} else if freeCount == lastFileCount {
				// blocks: Fill all the free space with the last fileID. line: Remove the free space from the beginning, remove the last fileID and freespace at the end, iterate.
				blocks = append(blocks, fillSlice(lastFileID, freeCount)...)
				line = line[1 : len(line)-2]

				isFile = !isFile
				fileID++
			} else if freeCount > lastFileCount {
				// blocks: Use some of the free space to fill all the last fileIDs with. line: Reduce the amount of free space at the beginning, remove the last fileID and freespace, do not iterate.
				blocks = append(blocks, fillSlice(lastFileID, lastFileCount)...)
				line[0] = toRune(freeCount - lastFileCount)
				line = line[:len(line)-2]
			}
		}
	}

	checksum := calculateChecksum(blocks)

	elapsed := time.Since(start)

	fmt.Println("Checksum:", checksum, "Time:", elapsed)

	// Part 2

	start = time.Now()

	line = []rune(lines[0])
	blocks = []int{}
	space := -1
	fileIDToCount := make(map[int]int)

	isFile = false
	fileID = -1
	for _, count := range line {
		isFile = !isFile
		if isFile {
			fileID++
		}

		countInt := toInt(count)
		if isFile {
			blocks = append(blocks, fillSlice(fileID, countInt)...)
			fileIDToCount[fileID] = countInt
		} else {
			blocks = append(blocks, fillSlice(space, countInt)...)
		}
	}

	maxBlockIndex := len(blocks) - 1

	blocks = moveBlocks(blocks, fileIDToCount, fileID, maxBlockIndex, space)

	checksum = calculateChecksum(blocks)

	elapsed = time.Since(start)

	fmt.Println("Checksum:", checksum, "Time:", elapsed)
}

func moveBlocks(blocks []int, fileIDToCount map[int]int, fileID int, maxBlockIndex int, space int) []int {
	if fileID == -1 {
		return blocks
	}

	fileIDIndex := slices.Index(blocks, fileID)
	fileIDCount := fileIDToCount[fileID]
	for i, block := range blocks {
		// Only check for spaces before the index of the current fileID
		if i >= fileIDIndex {
			break
		}

		if block == space {
			// Get number of free spaces
			spaceCount := 0
			for ii := i; ii <= maxBlockIndex; ii++ {
				if blocks[ii] == space {
					spaceCount++
				} else {
					break
				}
			}

			// See if the fileIDs fit in the space
			if fileIDCount <= spaceCount {
				// Move fileIDs into the space
				fileIDBlocks := fillSlice(fileID, fileIDCount)
				newSpaceBlocks := fillSlice(space, fileIDCount)
				remainingSpaceBlocks := fillSlice(space, spaceCount-fileIDCount)

				blocks = slices.Concat(
					blocks[:i],
					fileIDBlocks,
					remainingSpaceBlocks,
					blocks[i+spaceCount:fileIDIndex],
					newSpaceBlocks,
					blocks[fileIDIndex+fileIDCount:],
				)

				return moveBlocks(blocks, fileIDToCount, fileID-1, maxBlockIndex, space)
			}
		}
	}

	return moveBlocks(blocks, fileIDToCount, fileID-1, maxBlockIndex, space)
}

func calculateChecksum(blocks []int) int {
	checksum := 0
	for i, fileID := range blocks {
		if fileID == -1 {
			continue
		}
		checksum += i * fileID
	}
	return checksum
}

func fillSlice(x int, count int) []int {
	slice := make([]int, count)
	for i := range slice {
		slice[i] = x
	}
	return slice
}

func toInt(r rune) int {
	return int(r - '0')
}

func toRune(i int) rune {
	return rune(i + '0')
}
