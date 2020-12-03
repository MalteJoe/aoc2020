package main

import (
	"strings"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := strings.NewReader(
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	parsedInput, err := ReadInput(input)
	if err != nil {
		t.Fatalf(`Error in input: %v`, err)
	}

	output := CountTrees(parsedInput)

	if output != 7 {
		t.Fatalf(`expected 7, got %d`, output)
	}
}
