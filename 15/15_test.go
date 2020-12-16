package main

import (
	"testing"
)

// Tests
func TestPart1(t *testing.T) {
	assertIntEquals(t, 436, Part1([]int{0, 3, 6}))
	assertIntEquals(t, 1, Part1([]int{1, 3, 2}))
	assertIntEquals(t, 10, Part1([]int{2, 1, 3}))
	assertIntEquals(t, 27, Part1([]int{1, 2, 3}))
	assertIntEquals(t, 78, Part1([]int{2, 3, 1}))
	assertIntEquals(t, 438, Part1([]int{3, 2, 1}))
	assertIntEquals(t, 1836, Part1([]int{3, 1, 2}))
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
