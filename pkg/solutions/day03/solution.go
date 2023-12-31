package day03

import (
	"bufio"
	"fmt"
	"os"
	"utils"
)

const inputPath string = "inputs/day03.txt"

func findSharedGift(aGifts []rune, bGifts []rune) rune {
	for _, aGift := range aGifts {
		for _, bGift := range bGifts {
			if aGift == bGift {
				return aGift
			}
		}
	}

	panic("No shared gift has been found!")
}

func getGiftPriority(gift rune) int {
	giftValue := int(gift)

	if giftValue <= 90 { // A-Z
		return giftValue - 38
	} else { // a-z
		return giftValue - 96
	}
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	prioritySum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		gifts := []rune(line)
		giftCount := len(gifts)
		firstHalf := gifts[0 : giftCount/2]
		secondHalf := gifts[giftCount/2 : giftCount]

		sharedGift := findSharedGift(firstHalf, secondHalf)
		giftPriority := getGiftPriority(sharedGift)

		prioritySum += giftPriority
	}

	fmt.Print("Day 03 - Solution 01: ")
	fmt.Println(prioritySum)
	fmt.Print("Day 03 - Solution 02: ")
	fmt.Println("???")
}
