package main

import (
	"strings"
	"testing"
)

// Tests
func TestPart1(t *testing.T) {
	input := `939
7,13,x,x,59,x,31,19`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 295, Part1(parsed))
}

func TestPart2(t *testing.T) {
	input := `939
7,13,x,x,59,x,31,19`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 1068781, Part2(parsed))
}

func TestPart2Examples(t *testing.T) {
	x := -1
	assertIntEquals(t, 3417, Part2(Input{0, []int{17, x, 13, 19}}))
	assertIntEquals(t, 754018, Part2(Input{0, []int{67, 7, 59, 61}}))
	assertIntEquals(t, 779210, Part2(Input{0, []int{67, x, 7, 59, 61}}))
	assertIntEquals(t, 1261476, Part2(Input{0, []int{67, 7, x, 59, 61}}))
	assertIntEquals(t, 1202161486, Part2(Input{0, []int{1789, 37, 47, 1889}}))
}

// Helper
func assertTrue(t *testing.T, b bool) {
	t.Helper()
	if !b {
		t.Fatal("Expected true")
	}
}

func assertIntEquals(t *testing.T, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func assertStringEquals(t *testing.T, expected string, actual string) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected '%s', got '%s'", expected, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
