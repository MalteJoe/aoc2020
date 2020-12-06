package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func sum(slice []int) (sum int) {
	sum = 0
	for _, v := range slice {
		sum += v
	}
	return
}

func NumberOfTrueValues(maps []map[byte]bool) (count []int) {
	count = make([]int, len(maps))
	for i, m := range maps {
		count[i] = 0
		for _, b := range m {
			if b {
				count[i]++
			}
		}
	}
	return
}

func ReadInput(input io.Reader) (result []map[byte]bool, err error) {
	result = make([]map[byte]bool, 0)

	scanner := bufio.NewScanner(input)

	currentGroup := make(map[byte]bool)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			result = append(result, currentGroup)
			currentGroup = make(map[byte]bool)
		} else {
			for _, answer := range scanner.Bytes() {
				currentGroup[answer] = true
			}
		}
	}
	result = append(result, currentGroup)

	return
}

func ReadInput2(input io.Reader) (result []map[byte]bool, err error) {
	result = make([]map[byte]bool, 0)

	scanner := bufio.NewScanner(input)

	currentGroup := make(map[byte]bool)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			result = append(result, currentGroup)
			currentGroup = make(map[byte]bool)
		} else {
			if len(currentGroup) == 0 {
				for _, answer := range scanner.Bytes() {
					currentGroup[answer] = true
				}
			} else {
				for k := range currentGroup {
					if !strings.Contains(scanner.Text(), string(k)) {
						currentGroup[k] = false
					}
				}
			}
		}
	}
	result = append(result, currentGroup)

	return
}

// https://adventofcode.com/2020/day/6
func main() {
	log.SetPrefix("06: ")
	log.SetFlags(0)

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := ReadInput2(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Input: %v", input)

	answer := sum(NumberOfTrueValues(input))

	log.Printf("Answer: %v", answer)

}
