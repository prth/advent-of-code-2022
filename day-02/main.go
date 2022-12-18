package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	opponentMoveMap := make(map[string]int)
	opponentMoveMap["A"] = 1
	opponentMoveMap["B"] = 2
	opponentMoveMap["C"] = 3

	replyMoveMapByPart1Rules := make(map[string]int)
	replyMoveMapByPart1Rules["X"] = 1
	replyMoveMapByPart1Rules["Y"] = 2
	replyMoveMapByPart1Rules["Z"] = 3

	totalScorePart1 := 0
	totalScorePart2 := 0

	for _, entry := range input {
		opponent := entry[0]

		// rules based on part 1
		reply := entry[1]
		totalScorePart1 += replyMoveMapByPart1Rules[reply]

		if opponentMoveMap[opponent] == replyMoveMapByPart1Rules[reply] {
			totalScorePart1 += 3
		} else if opponentMoveMap[opponent] == 3 && replyMoveMapByPart1Rules[reply] == 1 {
			totalScorePart1 += 6
		} else if replyMoveMapByPart1Rules[reply]-opponentMoveMap[opponent] == 1 {
			totalScorePart1 += 6
		}

		// rules based on part 2
		expectedResult := entry[1]
		if expectedResult == "Y" {
			totalScorePart2 += opponentMoveMap[opponent]
			totalScorePart2 += 3
		} else if entry[1] == "Z" {
			if (opponentMoveMap[opponent]) == 3 {
				totalScorePart2 += 1
			} else {
				totalScorePart2 += opponentMoveMap[opponent] + 1
			}
			totalScorePart2 += 6
		} else {
			if (opponentMoveMap[opponent]) == 1 {
				totalScorePart2 += 3
			} else {
				totalScorePart2 += opponentMoveMap[opponent] - 1
			}
		}
	}

	answer1 = totalScorePart1
	answer2 = totalScorePart2

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() ([][]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Fields(line))
	}

	return input, scanner.Err()
}
