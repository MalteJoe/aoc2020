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
type Input []Instruction
type Instruction struct {
	Lvalue string
	RValue string
}

type state struct {
	mask0 uint64
	mask1 uint64
	mem   map[uint64]uint64
}

func ReadInput(input io.Reader) (result Input, err error) {
	scanner := bufio.NewScanner(input)
	result = make(Input, 0)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " = ")
		result = append(result, Instruction{parts[0], parts[1]})
	}

	err = scanner.Err()

	return
}

// Execute the initialization program.
// What is the sum of all values left in memory after it completes?
func Part1(input Input) (result int) {
	state := state{1<<36 - 1, 0, make(map[uint64]uint64)}
	for _, cmd := range input {
		if cmd.Lvalue == "mask" {
			state.mask0, _ = strconv.ParseUint(strings.ReplaceAll(cmd.RValue, "X", "1"), 2, 36)
			state.mask1, _ = strconv.ParseUint(strings.ReplaceAll(cmd.RValue, "X", "0"), 2, 36)
		} else {
			addr, _ := strconv.ParseUint(cmd.Lvalue[4:len(cmd.Lvalue)-1], 10, 36)
			value, _ := strconv.ParseUint(cmd.RValue, 10, 36)
			state.mem[addr] = value&state.mask0 | state.mask1
		}
	}
	for _, v := range state.mem {
		result += int(v)
	}
	return
}

// Execute the initialization program using an emulator for a version 2 decoder chip.
// What is the sum of all values left in memory after it completes?
func Part2(input Input) (result int) {
	mem := make(map[uint64]uint64)
	var mask string
	for _, cmd := range input {
		if cmd.Lvalue == "mask" {
			mask = cmd.RValue
		} else {
			addr, _ := strconv.ParseUint(cmd.Lvalue[4:len(cmd.Lvalue)-1], 10, 36)
			value, _ := strconv.ParseUint(cmd.RValue, 10, 36)
			for _, masked := range addresses(addr, mask) {
				mem[masked] = value
			}
		}
	}
	for _, v := range mem {
		result += int(v)
	}
	return
}

func addresses(addr uint64, mask string) (result []uint64) {
	if !strings.Contains(mask, "X") {
		single, _ := strconv.ParseUint(mask, 2, 36)
		return []uint64{addr | single}
	}
	result = addresses(addr, strings.Replace(mask, "X", "1", 1))
	result = append(result, addresses(addr&^(1<<(35-strings.Index(mask, "X"))), strings.Replace(mask, "X", "0", 1))...)
	return
}

// https://adventofcode.com/2020/day/14
func main() {
	log.SetPrefix("14: ")
	log.SetFlags(log.Ltime)

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
