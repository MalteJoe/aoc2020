package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func hasKey(m map[string]string, key string) bool {
	_, present := m[key]
	return present
}

func CountValidPassports(passports []map[string]string) (count int) {
	count = 0
	for _, passport := range passports {
		hasAllRequiredKeys := true
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "byr")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "iyr")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "eyr")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "hgt")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "hcl")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "ecl")
		hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "pid")
		// hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "cid")
		if hasAllRequiredKeys {
			count++
		}
	}

	return
}

func ReadInput(input io.Reader) (result []map[string]string, err error) {
	keyvalueRegex := regexp.MustCompile("^([^:]+):(.+)$")
	result = make([]map[string]string, 0)

	scanner := bufio.NewScanner(input)

	currentPassport := make(map[string]string)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, currentPassport)
			currentPassport = make(map[string]string)
		} else {
			reader := strings.NewReader(scanner.Text())
			entryScanner := bufio.NewScanner(reader)
			entryScanner.Split(bufio.ScanWords)
			for entryScanner.Scan() {
				match := keyvalueRegex.FindStringSubmatch(entryScanner.Text())
				if match == nil || len(match) != 3 {
					return nil, errors.New("Invalid line: " + scanner.Text())
				}
				currentPassport[match[1]] = match[2]
			}
		}
	}
	result = append(result, currentPassport)

	return
}

// https://adventofcode.com/2020/day/4
func main() {
	log.SetPrefix("04: ")
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

	answer := CountValidPassports(input)

	log.Printf("Answer: %d", answer)

}
