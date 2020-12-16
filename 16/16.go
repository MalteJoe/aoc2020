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
	rules         map[string][]Range
	myTicket      []int
	nearbyTickets [][]int
}

type Range struct {
	from, to int
}

func ReadInput(input io.Reader) (result Input, err error) {
	scanner := bufio.NewScanner(input)
	// rules
	result.rules = make(map[string][]Range)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			break
		}
		rule := scanner.Text()
		field := rule[:strings.Index(rule, ":")]
		ranges := strings.Split(rule[strings.Index(rule, ":")+2:], " or ")
		result.rules[field] = make([]Range, len(ranges))
		for i, v := range ranges {
			validRange := strings.Split(v, "-")
			result.rules[field][i] = Range{parseInt(validRange[0]), parseInt(validRange[1])}
		}
	}

	// my ticket
	scanner.Scan()
	scanner.Scan()
	values := strings.Split(scanner.Text(), ",")
	result.myTicket = make([]int, len(values))
	for i, v := range values {
		result.myTicket[i] = parseInt(v)
	}
	scanner.Scan()

	// nearby tickets
	scanner.Scan()
	result.nearbyTickets = make([][]int, 0)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		ticketValues := make([]int, len(values))
		result.nearbyTickets = append(result.nearbyTickets, ticketValues)
		for i, v := range values {
			ticketValues[i] = parseInt(v)
		}
	}

	err = scanner.Err()

	return
}

// Consider the validity of the nearby tickets you scanned.
// What is your ticket scanning error rate?
func Part1(input Input) (result int) {
	for _, ticket := range input.nearbyTickets {
		for _, value := range ticket {
			if !matchesAny(value, input.rules) {
				result += value
			}
		}
	}
	return
}

func matchesAny(value int, rules map[string][]Range) bool {
	for _, ranges := range rules {
		for _, ruleRange := range ranges {
			if value >= ruleRange.from && value <= ruleRange.to {
				return true
			}
		}
	}
	return false
}

// TODO
func Part2(input Input) (result int) {
	return -1
}

// https://adventofcode.com/2020/day/16
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
