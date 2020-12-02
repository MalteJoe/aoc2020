package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	MinOccurs int
	MaxOccurs int
	Character string
}

type passwordWithPolicy struct {
	policy   Policy
	password string
}

func PasswordValid(policy Policy, password string) bool {
	count := strings.Count(password, policy.Character)
	return count <= policy.MaxOccurs && count >= policy.MinOccurs
}

func countValidPasswords(inputs []passwordWithPolicy) uint {
	var count uint = 0
	for _, v := range inputs {
		if PasswordValid(v.policy, v.password) {
			count++
		}
	}
	return count
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
		minOccursString, err1 := strconv.ParseInt(match[1], 10, 64)
		maxOccursString, err2 := strconv.ParseInt(match[2], 10, 64)
		if match == nil || err1 != nil || err2 != nil {
			return nil, errors.New("Invalid line: " + line)
		}
		minOccurs := int(minOccursString)
		maxOccurs := int(maxOccursString)
		char := match[3]
		password := match[4]

		result = append(result, passwordWithPolicy{Policy{minOccurs, maxOccurs, char}, password})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
