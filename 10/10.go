package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

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

// What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?
func Solve1(input Input) (result int) {
	differences := make(map[int]int, 4)
	currJoltage := 0
	targetJoltage := max(input) + 3
	for currJoltage < targetJoltage {
		if targetJoltage-currJoltage <= 3 {
			differences[targetJoltage-currJoltage]++
			break
		}
		adapterFound := false
		for i := 1; i < 4; i++ {
			if contains(input, currJoltage+i) {
				adapterFound = true
				differences[i]++
				currJoltage += i
				break
			}
		}
		if !adapterFound {
			log.Fatal("No solution")
			return
		}
	}

	return differences[1] * differences[3]

}

func countArrangements(from int, adapters []int, knownArrangements map[int]int) (result int) {
	known, p := knownArrangements[from]
	if p {
		return known
	}
	if len(adapters) == 0 {
		return 1
	}
	for i := 1; i < 4; i++ {
		adapterIdx := indexOf(adapters, from+i)
		if adapterIdx != -1 {
			remainingAdapters := make([]int, 0)
			copy(remainingAdapters, adapters[:adapterIdx])
			remainingAdapters = append(remainingAdapters, adapters[adapterIdx+1:]...)
			result += countArrangements(from+i, remainingAdapters, knownArrangements)
		}
	}
	knownArrangements[from] = result
	return
}

// What is the total number of distinct ways you can arrange the adapters
// to connect the charging outlet to your device?
func Solve2(input Input) (result int) {
	sort.Ints(input)
	return countArrangements(0, input, make(map[int]int, 0))
}

// https://adventofcode.com/2020/day/10
func main() {
	log.SetPrefix("10: ")
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
