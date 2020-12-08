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

// Fix the program so that it terminates normally by
// changing exactly one jmp (to nop) or nop (to jmp).
// What is the value of the accumulator after the program terminates?
func Solve2(input Input) (result int) {
	acc := &result

	for corruptedIdx := 0; corruptedIdx < len(input); corruptedIdx++ {
		*acc = 0
		pc := 0
		visited := make(map[int]bool, 0)

		for !visited[pc] && pc < len(input) {
			visited[pc] = true
			op := input[pc].Op
			if corruptedIdx == pc {
				if op == "nop" {
					op = "jmp"
				} else if op == "jmp" {
					op = "nop"
				}
			}
			switch op {
			case "acc":
				*acc += input[pc].Arg
				pc++
			case "nop":
				pc++
			case "jmp":
				pc += input[pc].Arg
			}
		}
		if !visited[pc] {
			return
		}
	}

	log.Fatal("No Solution")
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
