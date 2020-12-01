package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FindProduct(input []uint64) uint64 {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				product := input[i] * input[j]
				log.Printf("Found match: %d + %d = 2020; Product: %d", input[i], input[j], product)
				return product
			}
		}
	}
	return 0
}

// https://adventofcode.com/2020/day/1
func main() {
	log.SetPrefix("01: ")
	log.SetFlags(0)

	input, err := readInput("./input")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(input)

	answer := FindProduct(input)

	log.Printf("Answer: %d", answer)

}

func readInput(path string) ([]uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]uint64, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
