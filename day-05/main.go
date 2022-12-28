package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	answer1 := processForCrateMover9000()
	answer2 := processForCrateMover9001()

	log.Printf("Answer #1 :: %s", answer1)
	log.Printf("Answer #2 :: %s", answer2)
}

func processForCrateMover9000() string {
	inputStackMap, moves, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		for i := 0; i < move.moveElementsCount; i++ {
			poppedElements := inputStackMap[move.fromStackNumber].pop(1)
			if len(poppedElements) == 1 {
				inputStackMap[move.toStackNumber].pushTop(poppedElements[0])
			}
		}
	}

	var topOfEachStack string

	for i := 1; i <= len(inputStackMap); i++ {
		topOfEachStack += inputStackMap[i].head()
	}

	return topOfEachStack
}

func processForCrateMover9001() string {
	inputStackMap, moves, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		poppedElements := inputStackMap[move.fromStackNumber].pop(move.moveElementsCount)
		inputStackMap[move.toStackNumber].pushTopAll(poppedElements)
	}

	var topOfEachStack string

	for i := 1; i <= len(inputStackMap); i++ {
		topOfEachStack += inputStackMap[i].head()
	}

	return topOfEachStack
}

func getInput() (map[int]*Stack, []Move, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputStackMap := make(map[int]*Stack)
	var moves []Move

	processStackInput := true
	processMoveInput := false

	for scanner.Scan() {
		line := scanner.Text()

		if processStackInput && strings.Contains(line, "[") {
			for i := 0; i < len(line)-1; i += 4 {
				stackNumber := i/4 + 1
				crateId := line[i+1 : i+2]

				if crateId != " " {
					if _, ok := inputStackMap[stackNumber]; !ok {
						inputStackMap[stackNumber] = &Stack{}
					}

					inputStackMap[stackNumber].pushBottom(crateId)
				}
			}
		} else {
			processStackInput = false
		}

		if processMoveInput || strings.Contains(line, "move") {
			processMoveInput = true
			words := strings.Fields(line)
			moveElementsCount, _ := strconv.Atoi(words[1])
			fromStackNumber, _ := strconv.Atoi(words[3])
			toStackNumber, _ := strconv.Atoi(words[5])

			move := Move{
				moveElementsCount: moveElementsCount,
				fromStackNumber:   fromStackNumber,
				toStackNumber:     toStackNumber,
			}

			moves = append(moves, move)
		}
	}

	return inputStackMap, moves, scanner.Err()
}
