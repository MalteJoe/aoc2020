package main

import (
	"testing"
)

func assertEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestDecodeSeat1(t *testing.T) {
	input := "FBFBBFFRLR"

	seat := DecodeSeat(input)

	assertEquals(t, 44, seat.Row)
	assertEquals(t, 5, seat.Column)
	assertEquals(t, 357, SeatId(seat))
}

func TestDecodeSeat2(t *testing.T) {
	input := "BFFFBBFRRR"

	seat := DecodeSeat(input)

	assertEquals(t, 70, seat.Row)
	assertEquals(t, 7, seat.Column)
	assertEquals(t, 567, SeatId(seat))
}

func TestDecodeSeat3(t *testing.T) {
	input := "FFFBBBFRRR"

	seat := DecodeSeat(input)

	assertEquals(t, 14, seat.Row)
	assertEquals(t, 7, seat.Column)
	assertEquals(t, 119, SeatId(seat))
}

func TestDecodeSeat4(t *testing.T) {
	input := "BBFFBBFRLL"

	seat := DecodeSeat(input)

	assertEquals(t, 102, seat.Row)
	assertEquals(t, 4, seat.Column)
	assertEquals(t, 820, SeatId(seat))
}
