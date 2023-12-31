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

func findGroupBadge(gifts1 []rune, gifts2 []rune, gifts3 []rune) rune {
	for _, gift1 := range gifts1 {
		for _, gift2 := range gifts2 {
			if gift1 != gift2 {
				continue
			}

			for _, gift3 := range gifts3 {
				if gift1 == gift3 {
					return gift1
				}
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

func getDuplicatedGiftPriority(gifts []rune) int {
	giftCount := len(gifts)
	firstHalf := gifts[0 : giftCount/2]
	secondHalf := gifts[giftCount/2 : giftCount]

	sharedGift := findSharedGift(firstHalf, secondHalf)

	return getGiftPriority(sharedGift)
}

func getGroupBadgePriority(gifts1 []rune, gifts2 []rune, gifts3 []rune) int {
	groupBadge := findGroupBadge(gifts1, gifts2, gifts3)
	badgePriority := getGiftPriority(groupBadge)

	return badgePriority
}

func read3Lines(fileScanner *bufio.Scanner) (string, string, string, bool) {
	if !fileScanner.Scan() {
		return "", "", "", false
	}

	line1 := fileScanner.Text()

	if !fileScanner.Scan() {
		return "", "", "", false
	}

	line2 := fileScanner.Text()

	if !fileScanner.Scan() {
		return "", "", "", false
	}

	line3 := fileScanner.Text()

	return line1, line2, line3, true
}

func Solve() {
	inputFile, err := os.Open(inputPath)

	utils.HandleError(err)
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	duplicatePrioritySum := 0
	badgePrioritySum := 0

	for {
		line1, line2, line3, hasMore := read3Lines(fileScanner)

		if !hasMore {
			break
		}

		gifts1 := []rune(line1)
		gifts2 := []rune(line2)
		gifts3 := []rune(line3)

		duplicatePrioritySum += getDuplicatedGiftPriority(gifts1)
		duplicatePrioritySum += getDuplicatedGiftPriority(gifts2)
		duplicatePrioritySum += getDuplicatedGiftPriority(gifts3)
		badgePrioritySum += getGroupBadgePriority(gifts1, gifts2, gifts3)
	}

	fmt.Print("Day 03 - Solution 01: ")
	fmt.Println(duplicatePrioritySum)
	fmt.Print("Day 03 - Solution 02: ")
	fmt.Println(badgePrioritySum)
}
