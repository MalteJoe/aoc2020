package main

import (
	"strings"
	"testing"
)

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

// Tests
func TestReadInput1(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags."

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)
	t.Logf("Parsed: %v", parsed)

	assertIntEquals(t, 1, len(parsed))
	v, present := parsed["light red"]

	assertTrue(t, present)
	assertIntEquals(t, 2, len(v))
	assertIntEquals(t, 1, v[0].Count)
	assertStringEquals(t, "bright white", v[0].BagColor)
	assertIntEquals(t, 2, v[1].Count)
	assertStringEquals(t, "muted yellow", v[1].BagColor)
}

func TestReadInput2(t *testing.T) {
	input := "faded blue bags contain no other bags."

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)
	t.Logf("Parsed: %v", parsed)

	assertIntEquals(t, 1, len(parsed))
	v, present := parsed["faded blue"]

	assertTrue(t, present)
	assertIntEquals(t, 0, len(v))
}

func TestSolve1(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)

	assertIntEquals(t, 4, Solve(parsed))
}
