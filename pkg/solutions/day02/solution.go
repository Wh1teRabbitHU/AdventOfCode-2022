package day02

import (
	"bufio"
	"fmt"
	"os"
	"utils"
)

const inputPath string = "inputs/day02.txt"

const (
	Rock         = "Rock"
	RockPoint    = 1
	Paper        = "Paper"
	PaperPoint   = 2
	Scissor      = "Scissor"
	ScissorPoint = 3
	LooseCode    = 'X'
	LoosePoint   = 0
	DrawCode     = 'Y'
	DrawPoint    = 3
	WinCode      = 'Z'
	WinPoint     = 6
)

var RockRunes = []rune{'A', 'X'}
var PaperRunes = []rune{'B', 'Y'}
var ScissorRunes = []rune{'C', 'Z'}

func getChoice(choiceRune rune) (string, int) {
	for _, rune := range RockRunes {
		if rune == choiceRune {
			return Rock, RockPoint
		}
	}

	for _, rune := range PaperRunes {
		if rune == choiceRune {
			return Paper, PaperPoint
		}
	}

	for _, rune := range ScissorRunes {
		if rune == choiceRune {
			return Scissor, ScissorPoint
		}
	}

	panic(0)
}

func getMatchPoint(myChoice string, elfChoice string) int {
	if myChoice == elfChoice {
		return DrawPoint // Draw
	}

	win1 := myChoice == Rock && elfChoice == Scissor
	win2 := myChoice == Paper && elfChoice == Rock
	win3 := myChoice == Scissor && elfChoice == Paper

	if win1 || win2 || win3 {
		return WinPoint // Win
	}

	return LoosePoint // Loose
}

func getPointsFirstTime(choiceRunes []rune) int {
	elfRune := choiceRunes[0]
	myRune := choiceRunes[2]

	myChoice, choicePoint := getChoice(myRune)
	elfChoice, _ := getChoice(elfRune)
	matchPoint := getMatchPoint(myChoice, elfChoice)

	return matchPoint + choicePoint
}

func getPointsSecondTime(choiceRunes []rune) int {
	matchPoint, choicePoint := 0, 1
	elfRune := choiceRunes[0]
	myRune := choiceRunes[2]

	elfChoice, elfPoint := getChoice(elfRune)

	if myRune == LooseCode {
		if elfChoice == Rock {
			choicePoint = ScissorPoint
		} else if elfChoice == Scissor {
			choicePoint = PaperPoint
		}

		matchPoint = LoosePoint // Not necessary, just for better readability
	} else if myRune == DrawCode {
		choicePoint = elfPoint
		matchPoint = DrawPoint
	} else if myRune == WinCode {
		if elfChoice == Rock {
			choicePoint = PaperPoint
		} else if elfChoice == Paper {
			choicePoint = ScissorPoint
		}

		matchPoint = WinPoint
	}

	return matchPoint + choicePoint
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	pointsFirstTime, pointsSecondTime := 0, 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		choiceRunes := []rune(line)

		pointsFirstTime += getPointsFirstTime(choiceRunes)
		pointsSecondTime += getPointsSecondTime(choiceRunes)
	}

	fmt.Print("Day 02 - Solution 01: ")
	fmt.Println(pointsFirstTime)
	fmt.Print("Day 02 - Solution 02: ")
	fmt.Println(pointsSecondTime)
}
