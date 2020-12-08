package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
type Instruction struct {
	Op  string
	Arg int
}

type Input []Instruction

func ReadInput(input io.Reader) (result Input, err error) {
	re := regexp.MustCompile("^(nop|acc|jmp) ([+-]\\d+)$")
	result = make([]Instruction, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		if match == nil || len(match) != 3 {
			return nil, fmt.Errorf("Invalid line: '%s'; Match: %v", scanner.Text(), match)
		}
		result = append(result, Instruction{match[1], parseInt(match[2])})
	}

	err = scanner.Err()

	return
}

// Run your copy of the boot code.
// Immediately before any instruction is executed a second time,
// what value is in the accumulator?
func Solve1(input Input) (result int) {
	acc := &result
	pc := 0
	visited := make(map[int]bool, 0)

	for !visited[pc] {
		visited[pc] = true
		switch input[pc].Op {
		case "acc":
			*acc += input[pc].Arg
			fallthrough
		case "nop":
			pc++
		case "jmp":
			pc += input[pc].Arg
		}

	}

	return
}

// TODO
func Solve2(input Input) (result int) {
	return
}

// https://adventofcode.com/2020/day/8
func main() {
	log.SetPrefix("08: ")
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
