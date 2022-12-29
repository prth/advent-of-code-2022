package main

import (
	"bufio"
	"log"
	"os"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	charSequenceCharMap := make(map[string]int)
	var firstStartOfPacketMarker int
	var firstStartOfMessageMarker int

	for index, c := range input {
		char := string(c)
		if firstLocationOfRepeatedChar, ok := charSequenceCharMap[char]; ok {
			for eachSeqChar, eachSeqCharLocation := range charSequenceCharMap {
				if eachSeqCharLocation <= firstLocationOfRepeatedChar {
					delete(charSequenceCharMap, eachSeqChar)
				} else {
					charSequenceCharMap[eachSeqChar] -= firstLocationOfRepeatedChar
				}
			}
		}

		charSequenceCharMap[char] = len(charSequenceCharMap) + 1

		if firstStartOfPacketMarker == 0 && len(charSequenceCharMap) == 4 {
			firstStartOfPacketMarker = index + 1
		}

		if firstStartOfMessageMarker == 0 && len(charSequenceCharMap) == 14 {
			firstStartOfMessageMarker = index + 1
		}

		if firstStartOfPacketMarker > 0 && firstStartOfMessageMarker > 0 {
			break
		}
	}

	answer1 := firstStartOfPacketMarker
	answer2 := firstStartOfMessageMarker

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getInput() (string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		return line, nil
	}

	return "", scanner.Err()
}
