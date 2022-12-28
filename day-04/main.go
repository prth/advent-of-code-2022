package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type Pair struct {
	sections []Section
}

type Section struct {
	startId int
	endId   int
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	completeOverlapCount := 0
	anyOverlapCount := 0

	for _, pair := range input {
		if pair.sections[0].endId < pair.sections[1].startId || pair.sections[0].startId > pair.sections[1].endId {
			continue
		}

		anyOverlapCount++

		if pair.sections[0].startId <= pair.sections[1].startId && pair.sections[0].endId >= pair.sections[1].endId {
			completeOverlapCount++
		} else if pair.sections[1].startId <= pair.sections[0].startId && pair.sections[1].endId >= pair.sections[0].endId {
			completeOverlapCount++
		}
	}

	answer1 := completeOverlapCount
	answer2 := anyOverlapCount

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() ([]Pair, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []Pair

	for scanner.Scan() {
		line := scanner.Text()

		var pair Pair

		for _, sectionStr := range strings.Split(line, ",") {
			sectionIds := strings.Split(sectionStr, "-")
			startId, _ := strconv.Atoi(sectionIds[0])
			endId, _ := strconv.Atoi(sectionIds[1])

			section := Section{
				startId: startId,
				endId:   endId,
			}

			pair.sections = append(pair.sections, section)
		}
		input = append(input, pair)
	}

	return input, scanner.Err()
}
