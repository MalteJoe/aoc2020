package main

import (
	"testing"
)

func TestPassword1(t *testing.T) {
	policy := Policy{1, 3, "a"}
	password := "abcde"

	output := PasswordValid(policy, password)

	if !output {
		t.Fatalf(`expected true, got false`)
	}
}

func TestPassword2(t *testing.T) {
	policy := Policy{1, 3, "b"}
	password := "cdefg"

	output := PasswordValid(policy, password)

	if output {
		t.Fatalf(`expected false, got true`)
	}
}

func TestPassword3(t *testing.T) {
	policy := Policy{2, 9, "c"}
	password := "ccccccccc"

	output := PasswordValid(policy, password)

	if !output {
		t.Fatalf(`expected true, got false`)
	}
}
