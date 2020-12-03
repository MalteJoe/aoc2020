package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Slope struct {
	X, Y int
}

func CountAndMultiplyForEachSlope(treemap [][]byte, slopes []Slope) int {
	product := 1
	for i := 0; i < len(slopes); i++ {
		product *= CountTrees(treemap, slopes[i])
	}
	return product
}

func CountTrees(treemap [][]byte, slope Slope) int {
	var count, x, y, i int = 0, 0, 0, 0
	for y < len(treemap) {
		if treemap[y][x] == '#' {
			count++
		}
		i++
		y = i * slope.Y
		x = i * slope.X % len(treemap[0])
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

	slopes := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}

	answer := CountAndMultiplyForEachSlope(input, slopes)

	log.Printf("Answer: %d", answer)

}
