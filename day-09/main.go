package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type Move struct {
	direction  Direction
	stepsCount int
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Position struct {
	row    int
	column int
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	ropeKnotMap := make(map[int]*Position)
	maxRopeKnotsLengthToCheck := 10

	for i := 0; i < maxRopeKnotsLengthToCheck; i++ {
		ropeKnotMap[i] = &Position{
			row:    0,
			column: 0,
		}
	}

	ropeKnotPositionTracker := make(map[int]map[string]int)
	knotsToTrackPosition := [2]int{1, 9}

	for _, knotNoToTrack := range knotsToTrackPosition {
		ropeKnotPositionTracker[knotNoToTrack] = make(map[string]int)
		defaultPositionKey := getPositionMapKey(Position{row: 0, column: 0})
		ropeKnotPositionTracker[knotNoToTrack][defaultPositionKey] = 1
	}

	for _, move := range input {
		for eachStep := 1; eachStep <= move.stepsCount; eachStep++ {
			switch move.direction {
			case Right:
				ropeKnotMap[0].column += 1
			case Left:
				ropeKnotMap[0].column -= 1
			case Up:
				ropeKnotMap[0].row += 1
			case Down:
				ropeKnotMap[0].row -= 1
			}

			knotNoToEvaluate := 1

			for knotNoToEvaluate < maxRopeKnotsLengthToCheck {
				relativeHeadKnot := ropeKnotMap[knotNoToEvaluate-1]
				relativeTailKnot := ropeKnotMap[knotNoToEvaluate]
				isRelativeTailKnotPositionChanged := false

				rowDifference := relativeHeadKnot.row - relativeTailKnot.row
				columnDifference := relativeHeadKnot.column - relativeTailKnot.column

				if math.Abs(float64(rowDifference)) == 2 {
					relativeTailKnot.row = relativeTailKnot.row + rowDifference/2

					if math.Abs(float64(columnDifference)) == 1 {
						relativeTailKnot.column = relativeHeadKnot.column
					}
					isRelativeTailKnotPositionChanged = true
				}

				if math.Abs(float64(columnDifference)) == 2 {
					relativeTailKnot.column = relativeTailKnot.column + columnDifference/2

					if math.Abs(float64(rowDifference)) == 1 {
						relativeTailKnot.row = relativeHeadKnot.row
					}
					isRelativeTailKnotPositionChanged = true
				}

				if isRelativeTailKnotPositionChanged {
					if _, ok := ropeKnotPositionTracker[knotNoToEvaluate]; ok {
						relativeTailKnotMapKey := getPositionMapKey(*relativeTailKnot)
						if _, ok := ropeKnotPositionTracker[knotNoToEvaluate][relativeTailKnotMapKey]; !ok {
							ropeKnotPositionTracker[knotNoToEvaluate][relativeTailKnotMapKey] = 0
						}

						ropeKnotPositionTracker[knotNoToEvaluate][relativeTailKnotMapKey]++
					}

					knotNoToEvaluate++
				} else {
					break
				}
			}
		}

	}

	answer1 = len(ropeKnotPositionTracker[1])
	answer2 = len(ropeKnotPositionTracker[9])

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getPositionMapKey(pos Position) string {
	return strconv.Itoa(pos.row) + ":" + strconv.Itoa(pos.column)
}

func getInput() ([]Move, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []Move

	for scanner.Scan() {
		line := scanner.Text()
		moveStrParts := strings.Fields(line)

		direction, err := parseMoveDirectionStrPart(moveStrParts[0])
		if err != nil {
			return nil, err
		}

		stepsCount, _ := strconv.Atoi(moveStrParts[1])

		input = append(input, Move{
			direction:  direction,
			stepsCount: stepsCount,
		})
	}

	return input, scanner.Err()
}

func parseMoveDirectionStrPart(directionKey string) (Direction, error) {
	switch directionKey {
	case "R":
		return Right, nil
	case "L":
		return Left, nil
	case "D":
		return Down, nil
	case "U":
		return Up, nil
	}

	return -1, errors.New("INVALID DIRECTION KEY")
}
