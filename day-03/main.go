package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	priorityForPart1 := 0
	priorityForPart2 := 0

	trackMapForCommonBadge := make(map[rune]int)
	groupSize := 3

	for index, rucksack := range input {
		if len(rucksack)%2 != 0 {
			log.Fatal("Invalid rucksack input | size is not even")
		}

		if index%groupSize == 0 {
			trackMapForCommonBadge = make(map[rune]int)
		}

		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		firstCompartmentItemsMap := make(map[rune]bool)

		for _, item := range firstCompartment {
			firstCompartmentItemsMap[item] = true

			currentGroupLineNumber := index%groupSize + 1
			priorityForPart2 += evaluatePriorityForBadge(item, groupSize, currentGroupLineNumber, trackMapForCommonBadge)
		}

		itemsPresentInBothMap := make(map[rune]bool)
		for _, item := range secondCompartment {
			if !itemsPresentInBothMap[item] && firstCompartmentItemsMap[item] {
				itemsPresentInBothMap[item] = true

				priorityForPart1 += getPriority(item)
			}

			currentGroupLineNumber := index%groupSize + 1
			priorityForPart2 += evaluatePriorityForBadge(item, groupSize, currentGroupLineNumber, trackMapForCommonBadge)
		}
	}

	answer1 = priorityForPart1
	answer2 = priorityForPart2

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getPriority(item rune) int {
	priority := 0

	if unicode.IsUpper(item) {
		priority += 26
	}

	return priority + int(unicode.ToLower(item)) - int('a') + 1
}

func evaluatePriorityForBadge(item rune, groupSize int, currentGroupLineNumber int, trackMapForCommonBadge map[rune]int) int {
	if trackMapForCommonBadge[item] != currentGroupLineNumber {
		if currentGroupLineNumber == 1 {
			trackMapForCommonBadge[item] = 1
		} else if currentGroupLineNumber-trackMapForCommonBadge[item] == 1 {
			trackMapForCommonBadge[item] = currentGroupLineNumber

			if currentGroupLineNumber == groupSize {
				return getPriority(item)
			}
		}
	}

	return 0
}

func getInput() ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input, scanner.Err()
}
