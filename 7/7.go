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

type BagRule struct {
	Count    int
	BagColor string
}

type Input map[string][]BagRule

func ReadInput(input io.Reader) (result Input, err error) {
	ruleRe := regexp.MustCompile("^(.+) bags contain (.+)\\.$")
	contentRe := regexp.MustCompile("(\\d+) ([^,]+) bags?")

	result = make(map[string][]BagRule, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ruleMatch := ruleRe.FindStringSubmatch(scanner.Text())
		if ruleMatch == nil || len(ruleMatch) != 3 {
			return nil, fmt.Errorf("Invalid Rule: %s; Match: %v", scanner.Text(), ruleMatch)
		}
		contentMatch := contentRe.FindAllStringSubmatch(ruleMatch[2], -1)

		contents := make([]BagRule, len(contentMatch))
		for i, rule := range contentMatch {
			count, err := strconv.ParseInt(rule[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("Invalid Rule: %s; Match: %v", scanner.Text(), ruleMatch)
			}
			contents[i] = BagRule{int(count), rule[2]}
		}

		result[ruleMatch[1]] = contents
	}

	return
}

func canContain(input Input, container string, search string) bool {
	for _, rule := range input[container] {
		if search == rule.BagColor || canContain(input, rule.BagColor, search) {
			return true
		}
	}
	return false

}

func Solve(input Input) (result int) {
	for bagColor := range input {
		if canContain(input, bagColor, "shiny gold") {
			result++
		}
	}
	return
}

// https://adventofcode.com/2020/day/7
func main() {
	log.SetPrefix("07: ")
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

	// log.Printf("Input: %v", input)

	answer := Solve(input)

	log.Printf("Answer: %v", answer)

}
