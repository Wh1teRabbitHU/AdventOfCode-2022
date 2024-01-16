package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"utils"
)

const inputPath string = "../inputs/day06.txt"

func findDistinctSection(dataStream *string, distinctLength int) int {
	for i := distinctLength; i < len(*dataStream); i++ {
		section := (*dataStream)[i-distinctLength : i]
		markerFound := true

		for j := 0; j < distinctLength; j++ {
			char := section[j : j+1]
			occurence := strings.Count(section, char)

			if occurence > 1 {
				markerFound = false
				break
			}
		}

		if markerFound {
			return i
		}
	}

	return -1
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	dataStream := fileScanner.Text()
	startOfPacket := findDistinctSection(&dataStream, 4)
	startOfMessage := findDistinctSection(&dataStream, 14)

	fmt.Print("Day 06 - Solution 01: ")
	fmt.Println(startOfPacket)
	fmt.Print("Day 06 - Solution 02: ")
	fmt.Println(startOfMessage)
}
