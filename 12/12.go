package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// Challenge
type Input []Action
type Action struct {
	Type  byte
	Value int
}
type Vector struct {
	X int
	Y int
}

const (
	N = 0
	E = 1
	S = 2
	W = 3
)

func ReadInput(input io.Reader) (result Input, err error) {
	result = make(Input, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		result = append(result, Action{scanner.Bytes()[0], parseInt(string(scanner.Bytes()[1:]))})
	}

	err = scanner.Err()

	return
}

func execute(action Action, pos *Vector, heading *int) {
	if action.Type == 'F' {
		execute(Action{[]byte{'N', 'E', 'S', 'W'}[*heading], action.Value}, pos, heading)
	} else {
		switch action.Type {
		case 'N':
			pos.Y += action.Value
		case 'S':
			pos.Y -= action.Value
		case 'E':
			pos.X += action.Value
		case 'W':
			pos.X -= action.Value
		case 'L':
			*heading = (*heading + 4 - (action.Value / 90)) % 4
		case 'R':
			*heading = (*heading + (action.Value / 90)) % 4
		}
	}
}

// What is the Manhattan distance between that location
// and the ship's starting position?
func Part1(input Input) (result int) {
	position := Vector{0, 0}
	heading := E
	for _, a := range input {
		execute(a, &position, &heading)
	}
	return abs(position.X) + abs(position.Y)
}

// TODO
func Part2(input Input) (result int) {
	return -1
}

// https://adventofcode.com/2020/day/11
func main() {
	log.SetPrefix("11: ")
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

	log.Printf("Input: %v", input)

	log.Printf("Answer Part 1: %v", Part1(input))
	log.Printf("Answer Part 2: %v", Part2(input))

}

// utils
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sum(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return
}

func contains(slice []int, value int) bool {
	return indexOf(slice, value) != -1
}

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func min(slice []int) (min int) {
	if len(slice) == 0 {
		log.Fatal("min of empty slice")
	}
	min = slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return
}

func max(slice []int) (max int) {
	if len(slice) == 0 {
		log.Fatal("max of empty slice")
	}
	max = slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return
}

func parseInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Not an int: %s", s)
	}
	return result
}
