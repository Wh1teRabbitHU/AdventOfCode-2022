package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"utils"
)

const inputPath string = "../inputs/day05.txt"

func initCargo(cargo *[][]string, initData *[]string) {
	for i := len(*initData) - 1; i >= 0; i-- {
		row := (*initData)[i]
		rowLength := int((len(row) + 1) / 4)

		if len(*cargo) == 0 {
			*cargo = make([][]string, rowLength)
		}

		for j := 0; j < rowLength; j++ {
			if len((*cargo)[j]) == 0 {
				(*cargo)[j] = make([]string, 0)
			}

			crate := row[j*4 : (j+1)*4-1]

			if crate == "   " {
				continue
			}

			(*cargo)[j] = append((*cargo)[j], crate)
		}
	}
}

func moveCrates(cargo *[][]string, rawData string) {
	commandParts := strings.Split(rawData, " ")
	pieceCount, err := strconv.Atoi(commandParts[1])

	utils.HandleError(err)

	from, err := strconv.Atoi(commandParts[3])
	from--

	utils.HandleError(err)

	to, err := strconv.Atoi(commandParts[5])
	to--

	utils.HandleError(err)

	pieces := (*cargo)[from][len((*cargo)[from])-pieceCount : len((*cargo)[from])]
	(*cargo)[from] = (*cargo)[from][:len((*cargo)[from])-pieceCount]

	for i := len(pieces) - 1; i >= 0; i-- {
		piece := pieces[i]

		(*cargo)[to] = append((*cargo)[to], piece)
	}
}

func gatherTopCrates(cargo *[][]string) string {
	topCrates := ""

	for _, stack := range *cargo {
		topCrates += stack[len(stack)-1][1:2]
	}

	return topCrates
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	cargo := make([][]string, 0)
	initData := make([]string, 0)
	initStage := true

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			initCargo(&cargo, &initData)
			initStage = false

			continue
		}

		if initStage {
			if strings.Contains(line, "[") {
				initData = append(initData, line)
			}
		} else {
			moveCrates(&cargo, line)
		}
	}

	topCrates := gatherTopCrates(&cargo)

	fmt.Print("Day 05 - Solution 01: ")
	fmt.Println(topCrates)
	fmt.Print("Day 05 - Solution 02: ")
	fmt.Println("???")
}
