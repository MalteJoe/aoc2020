package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := []uint64{1721, 979, 366, 299, 675, 1456}
	var want uint64 = 514579

	output := FindProduct(input)

	if want != output {
		t.Fatalf(`expected %d, got %d`, want, output)
	}
}

func TestExampleOfThree(t *testing.T) {
	input := []uint64{1721, 979, 366, 299, 675, 1456}
	var want uint64 = 241861950

	output := FindProductOfThree(input)

	if want != output {
		t.Fatalf(`expected %d, got %d`, want, output)
	}
}
