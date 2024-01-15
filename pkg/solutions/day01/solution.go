package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"utils"
)

const inputPath string = "../inputs/day01.txt"

type TopCalories struct {
	first  int64
	second int64
	third  int64
}

func (c *TopCalories) sum() int64 {
	return c.first + c.second + c.third
}

func assignNewCalorie(topCalories *TopCalories, newCalorie int64) {
	if topCalories.third > newCalorie {
		return
	}

	topCalories.third = newCalorie

	if topCalories.second < topCalories.third {
		temp := topCalories.second
		topCalories.second = topCalories.third
		topCalories.third = temp

		if topCalories.first < topCalories.second {
			temp := topCalories.first
			topCalories.first = topCalories.second
			topCalories.second = temp
		}
	}
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var currentCounter int64 = 0
	topCalories := TopCalories{first: 0, second: 0, third: 0}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			assignNewCalorie(&topCalories, currentCounter)

			currentCounter = 0
		} else {
			calorie, err := strconv.ParseInt(line, 0, 0)
			utils.HandleError(err)

			currentCounter += calorie
		}
	}

	fmt.Print("Day 01 - Solution 01: ")
	fmt.Println(topCalories.first)
	fmt.Print("Day 01 - Solution 02: ")
	fmt.Println(topCalories.sum())
}
