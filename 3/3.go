package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func CountTrees(treemap [][]byte) int {
	var count, x, y, i int = 0, 0, 0, 0
	for y < len(treemap) {
		if treemap[y][x] == '#' {
			count++
		}
		i++
		y = i
		x = i * 3 % len(treemap[0])
	}
	return count
}

func ReadInput(input io.Reader) ([][]byte, error) {
	result := make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// https://adventofcode.com/2020/day/3
func main() {
	log.SetPrefix("03: ")
	log.SetFlags(0)

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input, err := ReadInput(file)
	if err != nil {
		log.Fatal(err)
	}

	answer := CountTrees(input)

	log.Printf("Answer: %d", answer)

}
