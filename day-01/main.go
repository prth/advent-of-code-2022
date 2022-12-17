package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	numOfTopElvesToTrack := 3
	topElvesCaloriesTotal := make([]int, numOfTopElvesToTrack)

	tempElfCaloriesTotal := 0

	// iterate over each calorie input
	for i := 1; i < len(input)+1; i++ {
		// do comparison for top elves on line break or end of input
		if i == len(input) || input[i] == -1 {
			// re-evaluate top elves calories total based on defined number of top elves to track (i.e. 3 as per the question)
			for e := 0; e < numOfTopElvesToTrack; e++ {
				if tempElfCaloriesTotal > topElvesCaloriesTotal[e] {
					for r := numOfTopElvesToTrack - 1; r >= e+1; r-- {
						topElvesCaloriesTotal[r] = topElvesCaloriesTotal[r-1]
					}
					topElvesCaloriesTotal[e] = tempElfCaloriesTotal
					break
				}
			}

			tempElfCaloriesTotal = 0
		} else {
			tempElfCaloriesTotal += input[i]
		}
	}

	totalSumOfTopElvesCalories := 0
	for i := 0; i < len(topElvesCaloriesTotal); i++ {
		totalSumOfTopElvesCalories = totalSumOfTopElvesCalories + topElvesCaloriesTotal[i]
	}

	answer1 = topElvesCaloriesTotal[0]
	answer2 = totalSumOfTopElvesCalories

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() ([]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			input = append(input, -1)
		} else {
			num, _ := strconv.Atoi(line)
			input = append(input, num)
		}
	}

	return input, scanner.Err()
}
