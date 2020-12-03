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

	output := CountTrees(parsedInput, Slope{3, 1})

	if output != 7 {
		t.Fatalf(`expected 7, got %d`, output)
	}
}

func TestExamplePart2(t *testing.T) {
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

	slopes := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}

	parsedInput, err := ReadInput(input)
	if err != nil {
		t.Fatalf(`Error in input: %v`, err)
	}

	output := CountAndMultiplyForEachSlope(parsedInput, slopes)

	if output != 336 {
		t.Fatalf(`expected 336, got %d`, output)
	}
}
