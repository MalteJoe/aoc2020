package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// utils
func sum(slice []int) (sum int) {
	sum = 0
	for _, v := range slice {
		sum += v
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

// Challenge
type Input []int

func ReadInput(input io.Reader) (result Input, err error) {
	result = make([]int, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		result = append(result, parseInt(scanner.Text()))
	}

	err = scanner.Err()

	return
}

func isSumOfTwo(preamble []int, value int) bool {
	for i := range preamble {
		for j := i + 1; j < len(preamble); j++ {
			if value == preamble[i]+preamble[j] {
				return true
			}
		}
	}
	return false
}

// The first step of attacking the weakness in the XMAS data is to
// find the first number in the list (after the preamble)
// which is not the sum of two of the 25 numbers before it.
// What is the first number that does not have this property?
func Solve1(input Input, preambleLength int) (result int) {
	for i, v := range input[preambleLength:] {
		preamble := input[i : i+preambleLength]
		if !isSumOfTwo(preamble, v) {
			return v
		}
	}

	log.Fatal("No solution")
	return
}

// Fix the program so that it terminates normally by
// changing exactly one jmp (to nop) or nop (to jmp).
// What is the value of the accumulator after the program terminates?
func Solve2(input Input, preambleLength int) (result int) {
	return Solve1(input, preambleLength)
}

// https://adventofcode.com/2020/day/9
func main() {
	log.SetPrefix("09: ")
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

	log.Printf("Answer Part 1: %v", Solve1(input, 25))
	log.Printf("Answer Part 2: %v", Solve2(input, 25))

}
