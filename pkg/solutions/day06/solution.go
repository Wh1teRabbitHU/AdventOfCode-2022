package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"utils"
)

const inputPath string = "../inputs/day06.txt"

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	dataStream := fileScanner.Text()
	startOfPacket := 0

	for i := 4; i < len(dataStream); i++ {
		section := dataStream[i-4 : i]
		markerFound := true

		for j := 0; j < 4; j++ {
			char := section[j : j+1]
			occurence := strings.Count(section, char)

			if occurence > 1 {
				markerFound = false
				break
			}
		}

		if markerFound {
			startOfPacket = i
			break
		}
	}

	fmt.Print("Day 06 - Solution 01: ")
	fmt.Println(startOfPacket)
	fmt.Print("Day 06 - Solution 02: ")
	fmt.Println("???")
}
