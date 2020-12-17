package util

import (
	"log"
	"strconv"
)

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Sum(slice []int) (sum int) {
	for _, v := range slice {
		sum += v
	}
	return
}

func RemoveAt(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Contains(slice []int, value int) bool {
	return IndexOf(slice, value) != -1
}

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Min(slice []int) (min int) {
	if len(slice) == 0 {
		log.Fatal("min of empty slice")
	}
	min = slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return
}

func Max(slice []int) (max int) {
	if len(slice) == 0 {
		log.Fatal("max of empty slice")
	}
	max = slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return
}

func ParseInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Not an int: %s", s)
	}
	return result
}
