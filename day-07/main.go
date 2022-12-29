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
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	totalRootSize := computeTotalSizeOfDirectoryTree(*input)

	filteredDirectores := filterDirectoriesByMaxTotalSize(*input, 100000)

	var totalSizeOfFilteredDirectors int
	for dir := range filteredDirectores {
		totalSizeOfFilteredDirectors += dir.dirTotalSize
	}

	unusedSpace := 70000000 - totalRootSize
	sizeToBeDeleted := 30000000 - unusedSpace
	closestFileSizeToDelete := filterDirectoryByClosestSize(*input, sizeToBeDeleted)

	answer1 := totalSizeOfFilteredDirectors
	answer2 := closestFileSizeToDelete.dirTotalSize

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

type Tree struct {
	root *Node
}

type Node struct {
	name         string
	parent       *Node
	children     map[string]*Node
	fileSize     int
	dirTotalSize int
}

func computeTotalSizeOfDirectoryTree(tree Tree) int {
	totalRootSize := processDirectoryNodeTotalSize(tree.root)
	return totalRootSize
}

func processDirectoryNodeTotalSize(node *Node) int {
	for _, childNode := range node.children {
		if len(childNode.children) == 0 {
			node.dirTotalSize += childNode.fileSize
		} else {
			node.dirTotalSize += processDirectoryNodeTotalSize(childNode)
		}
	}

	return node.dirTotalSize
}

func filterDirectoriesByMaxTotalSize(tree Tree, maxTotalSize int) map[*Node]struct{} {
	filteredDirectoresSet := make(map[*Node]struct{})
	processFilterDirectories(*tree.root, filteredDirectoresSet, maxTotalSize)

	return filteredDirectoresSet
}

func processFilterDirectories(node Node, filteredDirectoresSet map[*Node]struct{}, maxTotalSize int) {
	for _, childNode := range node.children {
		if len(childNode.children) != 0 {
			if childNode.dirTotalSize <= maxTotalSize {
				if _, ok := filteredDirectoresSet[childNode]; !ok {
					filteredDirectoresSet[childNode] = struct{}{}
				}
			}

			processFilterDirectories(*childNode, filteredDirectoresSet, maxTotalSize)
		}
	}
}

func filterDirectoryByClosestSize(tree Tree, closestSize int) Node {
	filteredDirectory := *tree.root
	processfilterDirectoryByClosestSize(tree.root, &filteredDirectory, closestSize)

	return filteredDirectory
}

func processfilterDirectoryByClosestSize(node *Node, filteredDirectory *Node, closestSize int) {
	for _, childNode := range node.children {
		if len(childNode.children) != 0 {
			if childNode.dirTotalSize >= closestSize && childNode.dirTotalSize < filteredDirectory.dirTotalSize {
				*filteredDirectory = *childNode
			}

			processfilterDirectoryByClosestSize(childNode, filteredDirectory, closestSize)
		}
	}
}

func getInput() (*Tree, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var directoryTree *Tree

	var tempDirectoryNode *Node

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$ cd") {
			currentDir := line[len("$ cd "):]

			if currentDir == ".." {
				tempDirectoryNode = tempDirectoryNode.parent
				continue
			} else if currentDir == "/" {
				if directoryTree != nil {
					tempDirectoryNode = directoryTree.root
				} else {
					tempDirectoryNode = &Node{
						name:     currentDir,
						children: make(map[string]*Node),
					}

					directoryTree = &Tree{
						root: tempDirectoryNode,
					}
				}

				continue
			} else {
				tempDirectoryNode = tempDirectoryNode.children[currentDir]
			}

		} else if !strings.HasPrefix(line, "$ ls") {

			if strings.HasPrefix(line, "dir") {
				childDirName := line[len("dir "):]
				if _, ok := tempDirectoryNode.children[childDirName]; !ok {
					tempDirectoryNode.children[childDirName] = &Node{
						name:     childDirName,
						children: make(map[string]*Node),
						parent:   tempDirectoryNode,
					}
				}
			} else {
				words := strings.Fields(line)
				childFileName := words[1]
				fileSize, _ := strconv.Atoi(words[0])
				if _, ok := tempDirectoryNode.children[childFileName]; !ok {
					tempDirectoryNode.children[childFileName] = &Node{
						name:     childFileName,
						fileSize: fileSize,
					}
				}
			}

		}
	}

	return directoryTree, scanner.Err()
}
