package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// Challenge
type Input [][]byte

func ReadInput(input io.Reader) (result Input, err error) {
	result = make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		result = append(result, line)
	}

	err = scanner.Err()

	return
}

func printMap(input Input) {
	for _, row := range input {
		log.Println(string(row))
	}
}

func adjacentSeats(input Input, row, col int) (output Input) {
	output = make([][]byte, 0)
	firstRow := max([]int{0, row - 1})
	lastRow := min([]int{len(input) - 1, row + 1})
	firstCol := max([]int{0, col - 1})
	lastCol := min([]int{len(input[0]) - 1, col + 1})
	rowSlice := input[firstRow : lastRow+1]
	for i := range rowSlice {
		output = append(output, make([]byte, lastCol+1-firstCol))
		copy(output[i], input[firstRow+i][firstCol:min([]int{lastCol + 1, len(input[firstRow+i])})])
	}
	centerSeatRow := max([]int{0, min([]int{len(input), row})}) - firstRow
	centerSeatCol := max([]int{0, min([]int{len(input[0]), col})}) - firstCol
	output[centerSeatRow][centerSeatCol] = '.'
	return
}

func round(input *Input) (seatsChanged bool) {
	output := make([][]byte, len(*input))
	for row := range *input {
		output[row] = make([]byte, len((*input)[row]))
		for col := range (*input)[row] {
			output[row][col] = (*input)[row][col]
			switch (*input)[row][col] {
			case 'L':
				if countOccupied(adjacentSeats(*input, row, col)) == 0 {
					seatsChanged = true
					output[row][col] = '#'
				}
			case '#':
				if countOccupied(adjacentSeats(*input, row, col)) >= 4 {
					seatsChanged = true
					output[row][col] = 'L'
				}
			}
		}
	}

	printMap(output)

	*input = output
	return
}

func countOccupied(input Input) (result int) {
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '#' {
				result++
			}
		}
	}
	return
}

// Simulate your seating area by applying the seating rules repeatedly
// until no seats change state. How many seats end up occupied?
func Solve1(input Input) (result int) {
	for round(&input) {
	}
	return countOccupied(input)
}

// What is the total number of distinct ways you can arrange the adapters
// to connect the charging outlet to your device?
func Solve2(input Input) (result int) {
	return Solve1(input)
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

	log.Printf("Answer Part 1: %v", Solve1(input))
	log.Printf("Answer Part 2: %v", Solve2(input))

}

// utils
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
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Not an int: %s", s)
	}
	return int(result)
}
