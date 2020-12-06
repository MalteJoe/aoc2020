package main

import (
	"strings"
	"testing"
)

func assertEquals(t *testing.T, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestSingleGroup(t *testing.T) {
	input := `abcx
abcy
abcz`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)
	t.Logf("Parsed: %v", parsed)

	assertEquals(t, 1, len(parsed))
	assertEquals(t, 6, len(parsed[0]))
}

func TestMultipleGroups(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`

	parsed, err := ReadInput(strings.NewReader(input))
	assertNoError(t, err)
	t.Logf("Parsed: %v", parsed)

	assertEquals(t, 5, len(parsed))
	assertEquals(t, 3, len(parsed[0]))
	assertEquals(t, 3, len(parsed[1]))
	assertEquals(t, 3, len(parsed[2]))
	assertEquals(t, 1, len(parsed[3]))
	assertEquals(t, 1, len(parsed[4]))
}
