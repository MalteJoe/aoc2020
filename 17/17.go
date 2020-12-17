package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input [][]byte

type Cube struct {
	X, Y, Z int
}

type void struct{}

type PocketDim map[Cube]void

func ReadInput(input io.Reader) (result Input, err error) {
	scanner := bufio.NewScanner(input)
	// rules
	result = make(Input, 0)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		result = append(result, line)
	}

	err = scanner.Err()

	return
}

// Starting with your given initial configuration, simulate six cycles.
// How many cubes are left in the active state after the sixth cycle?
func Part1(input Input) (result int) {
	pocketDim := make(PocketDim)
	for x := range input {
		for y := range input[x] {
			if input[x][y] == '#' {
				pocketDim[Cube{x, y, 0}] = void{}
			}
		}
	}

	for i := 1; i <= 6; i++ {
		pocketDim = simulate(pocketDim)
	}
	return len(pocketDim)
}

func simulate(dim PocketDim) PocketDim {
	activeNeighbours := make(map[Cube]byte)
	for cube := range dim {
		for x := cube.X - 1; x <= cube.X+1; x++ {
			for y := cube.Y - 1; y <= cube.Y+1; y++ {
				for z := cube.Z - 1; z <= cube.Z+1; z++ {
					if x == cube.X && y == cube.Y && z == cube.Z {
						continue
					}
					activeNeighbours[Cube{x, y, z}]++
				}
			}
		}
	}

	newDim := make(PocketDim)
	for cube, activeNeighbourCount := range activeNeighbours {
		_, active := dim[cube]
		if active && activeNeighbourCount >= 2 && activeNeighbourCount <= 3 ||
			!active && activeNeighbourCount == 3 {
			newDim[cube] = void{}
		}
	}

	return newDim
}

// TODO
func Part2(input Input) (result int) {
	return -1
}

// https://adventofcode.com/2020/day/17
func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := ReadInput(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Input: %v", input)

	log.Printf("Answer Part 1: %v", Part1(input))
	log.Printf("Answer Part 2: %v", Part2(input))
}
