package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func hasKey(m map[string]string, key string) bool {
	_, present := m[key]
	return present
}

func isNumericInRange(input string, min int64, max int64) bool {
	parsedInt, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Printf("Not a number: %s", input)
		return false
	}
	return parsedInt >= min && parsedInt <= max
}

func isValidHeight(input string) bool {
	re := regexp.MustCompile("^(\\d+)(cm|in)$")
	match := re.FindStringSubmatch(input)
	if match == nil || len(match) != 3 {
		log.Printf("Invalid Height: %s", input)
		return false
	}
	value, _ := strconv.ParseInt(match[1], 10, 64)
	if match[2] == "cm" {
		return value >= 150 && value <= 193
	}
	return value >= 59 && value <= 76
}

func matchRegex(input string, regex string) bool {
	match, err := regexp.MatchString(regex, input)
	if err != nil {
		log.Printf("Regex Error: %v, Regex: %s", err, regex)
		return false
	}
	return match
}

func CountValidPassportsWithValidation(passports []map[string]string) (count int) {
	count = 0
	for _, passport := range passports {
		isValid := true
		isValid = isValid && isNumericInRange(passport["byr"], 1920, 2002)
		isValid = isValid && isNumericInRange(passport["iyr"], 2010, 2020)
		isValid = isValid && isNumericInRange(passport["eyr"], 2020, 2030)
		isValid = isValid && isValidHeight(passport["hgt"])
		isValid = isValid && matchRegex(passport["hcl"], "^#[0-9a-f]{6}$")
		isValid = isValid && matchRegex(passport["ecl"], "^amb|blu|brn|gry|grn|hzl|oth$")
		isValid = isValid && matchRegex(passport["pid"], "^\\d{9}$")
		// hasAllRequiredKeys = hasAllRequiredKeys && hasKey(passport, "cid")
		if isValid {
			count++
		}
	}

	return
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

	answer := CountValidPassportsWithValidation(input)

	log.Printf("Answer: %d", answer)

}
