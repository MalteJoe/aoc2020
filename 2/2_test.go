package main

import (
	"testing"
)

func TestPassword1(t *testing.T) {
	policy := Policy{1, 3, 'a'}
	password := "abcde"

	output := PasswordValid(policy, password)

	if !output {
		t.Fatalf(`expected true, got false`)
	}
}
