package main

import (
	"strings"
	"testing"
)

// Tests
func TestSolve1(t *testing.T) {
	input := `F10
N3
F7
R90
F11`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 25, Part1(parsed))
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
