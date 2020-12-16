package main

import (
	"log"
	"strconv"
)

// Challenge
type Input []int

// Given your starting numbers, what will be the 2020th number spoken?
func Part1(input Input) (result int) {
	spokenNumbers := make(map[int]int)
	for turn, startingNumber := range input[:len(input)-1] {
		spokenNumbers[startingNumber] = turn + 1
	}

	lastNumber := input[len(input)-1]
	for turn := len(input); turn < 2020; turn++ {
		lastSpoken, wasSpoken := spokenNumbers[lastNumber]
		spokenNumbers[lastNumber] = turn
		if !wasSpoken {
			lastNumber = 0
		} else {
			lastNumber = turn - lastSpoken
		}
	}
	return lastNumber
}

// Execute the initialization program using an emulator for a version 2 decoder chip.
// What is the sum of all values left in memory after it completes?
func Part2(input Input) (result int) {
	return
}

// https://adventofcode.com/2020/day/15
func main() {
	log.Printf("Answer Part 1: %v", Part1([]int{2, 0, 6, 12, 1, 3}))
	//log.Printf("Answer Part 2: %v", Part2(input))
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
