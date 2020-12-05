package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Seat struct {
	Row, Column int
}

func DecodeSeat(seatRepr string) Seat {
	re := regexp.MustCompile("^([FB]{7})([LR]{3})$")
	match := re.FindStringSubmatch(seatRepr)
	if match == nil || len(match) != 3 {
		log.Fatalf("Invalid input: %s; match: %v", seatRepr, match)
	}

	return Seat{decode("F0B1", match[1]), decode("L0R1", match[2])}
}

func SeatId(seat Seat) int {
	return seat.Row*8 + seat.Column
}

func decode(replacement string, rowRepr string) int {
	replacer := strings.NewReplacer(strings.Split(replacement, "")...)
	binaryString := replacer.Replace(rowRepr)
	row, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(row)
}

func ReadInput(input io.Reader) (result []Seat, err error) {
	result = make([]Seat, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		result = append(result, DecodeSeat(scanner.Text()))
	}

	err = scanner.Err()

	return
}

// Part 1
func FindMaxId(seats []Seat) (max int) {
	max = -1
	for _, s := range seats {
		sID := SeatId(s)
		if sID > max {
			max = sID
		}
	}
	return
}

func contains(seats []Seat, seatID int) (contained bool) {
	contained = false
	for _, s := range seats {
		if SeatId(s) == seatID {
			contained = true
			break
		}
	}
	return
}

// Part 2
func FindMissingSeatId(seats []Seat) int {
	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			i := SeatId(Seat{row, col})

			if !contains(seats, i) && contains(seats, i-1) && contains(seats, i+1) {
				return i
			}
		}
	}
	return -1
}

// https://adventofcode.com/2020/day/5
func main() {
	log.SetPrefix("05: ")
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

	answer := FindMissingSeatId(input)

	log.Printf("Answer: %d", answer)

}
