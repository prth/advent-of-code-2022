package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type Tree struct {
	left   int
	right  int
	top    int
	bottom int
}

type Direction int

const (
	Left Direction = iota
	Right
	Top
	Bottom
)

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	treeShadeMap := computeTreeShadeMap(input)

	countInvisibleTrees := 0

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			tree := treeShadeMap[getCoordinatesMapKey(i, j)]
			treeHeight := input[i][j]

			if tree.left >= treeHeight && tree.right >= treeHeight && tree.top >= treeHeight && tree.bottom >= treeHeight {
				countInvisibleTrees++
			}
		}
	}

	maxScenicScore := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			tree := treeShadeMap[getCoordinatesMapKey(i, j)]
			treeHeight := input[i][j]

			scenicScore := 1

			if tree.left < treeHeight {
				scenicScore *= j
			} else {
				scenicScore *= countVisibleTrees(i, j, input, Left)
			}

			if tree.right < treeHeight {
				scenicScore *= len(input) - 1 - j
			} else {
				scenicScore *= countVisibleTrees(i, j, input, Right)
			}

			if tree.top < treeHeight {
				scenicScore *= i
			} else {
				scenicScore *= countVisibleTrees(i, j, input, Top)
			}

			if tree.bottom < treeHeight {
				scenicScore *= len(input) - 1 - i
			} else {
				scenicScore *= countVisibleTrees(i, j, input, Bottom)
			}

			if maxScenicScore < scenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	totalTreeCount := len(input) * len(input)

	answer1 := totalTreeCount - countInvisibleTrees
	answer2 := maxScenicScore

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func computeTreeShadeMap(trees [][]int) map[string]*Tree {
	treeShadeMap := make(map[string]*Tree)

	for i := 1; i < len(trees)-1; i++ {
		maxShadeLeft := trees[i][0]
		maxShadeRight := trees[i][len(trees)-1]

		maxShadeTop := trees[0][i]
		maxShadeBottom := trees[len(trees)-1][i]

		for j := 1; j < len(trees)-1; j++ {
			treeCoordinatesForLeftItr := getCoordinatesMapKey(i, j)
			treeHeightForLeftItr := trees[i][j]

			if _, ok := treeShadeMap[treeCoordinatesForLeftItr]; !ok {
				treeShadeMap[treeCoordinatesForLeftItr] = &Tree{
					left: maxShadeLeft,
				}
			} else {
				treeShadeMap[treeCoordinatesForLeftItr].left = maxShadeLeft
			}

			if maxShadeLeft < treeHeightForLeftItr {
				maxShadeLeft = treeHeightForLeftItr
			}

			treeCoordinatesForRightItr := getCoordinatesMapKey(i, len(trees)-1-j)
			treeHeightForRightItr := trees[i][len(trees)-1-j]

			if _, ok := treeShadeMap[treeCoordinatesForRightItr]; !ok {
				treeShadeMap[treeCoordinatesForRightItr] = &Tree{
					right: maxShadeRight,
				}
			} else {
				treeShadeMap[treeCoordinatesForRightItr].right = maxShadeRight
			}

			if maxShadeRight < treeHeightForRightItr {
				maxShadeRight = treeHeightForRightItr
			}

			treeCoordinatesForTopItr := getCoordinatesMapKey(j, i)
			treeHeightForTopItr := trees[j][i]

			if _, ok := treeShadeMap[treeCoordinatesForTopItr]; !ok {
				treeShadeMap[treeCoordinatesForTopItr] = &Tree{
					top: maxShadeTop,
				}
			} else {
				treeShadeMap[treeCoordinatesForTopItr].top = maxShadeTop
			}

			if maxShadeTop < treeHeightForTopItr {
				maxShadeTop = treeHeightForTopItr
			}

			treeCoordinatesForBottomItr := getCoordinatesMapKey(len(trees)-1-j, i)
			treeHeightForBottomItr := trees[len(trees)-1-j][i]

			if _, ok := treeShadeMap[treeCoordinatesForBottomItr]; !ok {
				treeShadeMap[treeCoordinatesForBottomItr] = &Tree{
					bottom: maxShadeBottom,
				}
			} else {
				treeShadeMap[treeCoordinatesForBottomItr].bottom = maxShadeBottom
			}

			if maxShadeBottom < treeHeightForBottomItr {
				maxShadeBottom = treeHeightForBottomItr
			}
		}
	}

	return treeShadeMap
}

func countVisibleTrees(treeRow int, treeColumn int, trees [][]int, direction Direction) int {
	count := 0

	treeHeight := trees[treeRow][treeColumn]

	switch direction {
	case Left:
		for i := treeColumn - 1; i >= 0; i-- {
			count++
			if trees[treeRow][i] >= treeHeight {
				break
			}
		}
	case Right:
		for i := treeColumn + 1; i < len(trees); i++ {
			count++
			if trees[treeRow][i] >= treeHeight {
				break
			}
		}
	case Top:
		for i := treeRow - 1; i >= 0; i-- {
			count++
			if trees[i][treeColumn] >= treeHeight {
				break
			}
		}
	case Bottom:
		for i := treeRow - 1; i < len(trees); i++ {
			count++
			if trees[i][treeColumn] >= treeHeight {
				break
			}
		}
	}

	return count
}

func getInput() ([][]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]int

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, convertStringToIntArray(line))
	}

	return input, scanner.Err()
}

func convertStringToIntArray(line string) []int {
	strs := strings.Split(line, "")
	ary := make([]int, len(strs))

	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}

	return ary
}

func getCoordinatesMapKey(x int, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}
