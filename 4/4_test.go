package main

import (
	"strings"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := strings.NewReader(
		`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)

	parsedInput, err := ReadInput(input)
	if err != nil {
		t.Fatalf(`Error in input: %v`, err)
	}

	if len(parsedInput) != 4 {
		t.Fatalf(`expected 4 inputs, got %d`, len(parsedInput))
	}

	t.Logf("Input: %v", parsedInput)

	output := CountValidPassports(parsedInput)

	if output != 2 {
		t.Fatalf(`expected 2, got %d`, output)
	}
}

func TestExamplePart2InvalidPassports(t *testing.T) {
	input := strings.NewReader(
		`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`)

	parsedInput, err := ReadInput(input)
	if err != nil {
		t.Fatalf(`Error in input: %v`, err)
	}

	if len(parsedInput) != 4 {
		t.Fatalf(`expected 4 inputs, got %d`, len(parsedInput))
	}

	t.Logf("Input: %v", parsedInput)

	output := CountValidPassportsWithValidation(parsedInput)

	if output != 0 {
		t.Fatalf(`expected 0, got %d`, output)
	}
}

func TestExamplePart2ValidPassports(t *testing.T) {
	input := strings.NewReader(
		`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`)

	parsedInput, err := ReadInput(input)
	if err != nil {
		t.Fatalf(`Error in input: %v`, err)
	}

	if len(parsedInput) != 4 {
		t.Fatalf(`expected 4 inputs, got %d`, len(parsedInput))
	}

	t.Logf("Input: %v", parsedInput)

	output := CountValidPassportsWithValidation(parsedInput)

	if output != 4 {
		t.Fatalf(`expected 4, got %d`, output)
	}
}
