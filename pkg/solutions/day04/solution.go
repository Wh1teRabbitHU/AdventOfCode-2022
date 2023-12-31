package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"utils"
)

const inputPath string = "inputs/day04.txt"

type AssignmentRange struct {
	start int
	end   int
}

func (r *AssignmentRange) partlyOverlap(other *AssignmentRange) bool {
	if r.start <= other.end && r.end >= other.start {
		return true
	}

	return false
}

func (r *AssignmentRange) fullyOverlap(other *AssignmentRange) bool {
	if r.start >= other.start && r.end <= other.end {
		return true
	} else if other.start >= r.start && other.end <= r.end {
		return true
	}

	return false
}

func resolveRanges(rangePairs []string) (AssignmentRange, AssignmentRange) {
	firstBoundaries := strings.Split(rangePairs[0], "-")
	firstStart, err1 := strconv.Atoi(firstBoundaries[0])
	firstEnd, err2 := strconv.Atoi(firstBoundaries[1])

	utils.HandleError(err1)
	utils.HandleError(err2)

	secondBoundaries := strings.Split(rangePairs[1], "-")
	secondStart, err1 := strconv.Atoi(secondBoundaries[0])
	secondEnd, err2 := strconv.Atoi(secondBoundaries[1])

	utils.HandleError(err1)
	utils.HandleError(err2)

	firstRange := AssignmentRange{start: firstStart, end: firstEnd}
	secondRange := AssignmentRange{start: secondStart, end: secondEnd}

	return firstRange, secondRange
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	partlyOverlapCounter := 0
	fullyOverlapCounter := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		rangePairs := strings.Split(line, ",")
		range1, range2 := resolveRanges(rangePairs)

		if range1.partlyOverlap(&range2) {
			partlyOverlapCounter++
		}

		if range1.fullyOverlap(&range2) {
			fullyOverlapCounter++
		}
	}

	fmt.Print("Day 04 - Solution 01: ")
	fmt.Println(fullyOverlapCounter)
	fmt.Print("Day 04 - Solution 02: ")
	fmt.Println(partlyOverlapCounter)
}
