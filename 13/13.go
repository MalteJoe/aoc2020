package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Challenge
type Input struct {
	Earliest int
	BusIDs   []int
}

func ReadInput(input io.Reader) (result Input, err error) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	result.Earliest = parseInt(scanner.Text())
	scanner.Scan()
	busIds := strings.Split(scanner.Text(), ",")
	result.BusIDs = make([]int, len(busIds))
	for i, busID := range busIds {
		if busID == "x" {
			result.BusIDs[i] = -1
		} else {
			result.BusIDs[i] = parseInt(busID)
		}
	}

	err = scanner.Err()

	return
}

// What is the ID of the earliest bus you can take to the airport
// multiplied by the number of minutes you'll need to wait for that bus?
func Part1(input Input) (result int) {
	bestBus := -1
	waitTimes := make(map[int]int)
	for _, busID := range input.BusIDs {
		if busID == -1 {
			continue
		}
		missedBy := input.Earliest % busID
		nextDeparture := input.Earliest
		if missedBy != 0 {
			nextDeparture += busID - missedBy
		}
		waitTimes[busID] = nextDeparture - input.Earliest
		if bestBus == -1 || waitTimes[busID] < waitTimes[bestBus] {
			bestBus = busID
		}
	}

	return bestBus * waitTimes[bestBus]
}

func Part2(input Input) (result int) {
	return -1
}

// https://adventofcode.com/2020/day/13
func main() {
	log.SetPrefix("13: ")
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
