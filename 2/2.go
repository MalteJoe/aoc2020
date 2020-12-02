package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Policy struct {
	MinOccurs uint
	MaxOccurs uint
	Character byte
}

type passwordWithPolicy struct {
	policy   Policy
	password string
}

func PasswordValid(policy Policy, password string) bool {
	return false
}

func countValidPasswords([]passwordWithPolicy) uint {
	return 0
}

// https://adventofcode.com/2020/day/2
func main() {
	log.SetPrefix("01: ")
	log.SetFlags(0)

	input, err := readInput("./input")
	if err != nil {
		log.Fatal(err)
	}

	answer := countValidPasswords(input)

	log.Printf("Answer: %d", answer)

}

func readInput(path string) ([]passwordWithPolicy, error) {
	re := regexp.MustCompile("(\\d+)-(\\d+) (.): (.+)")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([]passwordWithPolicy, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		minOccursString, err1 := strconv.ParseUint(match[1], 10, 64)
		maxOccursString, err2 := strconv.ParseUint(match[2], 10, 64)
		if match == nil || err1 != nil || err2 != nil {
			return nil, errors.New("Invalid line: " + line)
		}
		minOccurs := uint(minOccursString)
		maxOccurs := uint(maxOccursString)
		char := match[3][0]
		password := match[4]

		result = append(result, passwordWithPolicy{Policy{minOccurs, maxOccurs, char}, password})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
