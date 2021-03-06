package main

import (
	"strings"
	"testing"
)

// Tests
func TestPart1(t *testing.T) {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 165, Part1(parsed))
}

// Tests
func TestPart2(t *testing.T) {
	input := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 208, Part2(parsed))
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
