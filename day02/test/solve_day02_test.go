package testing

import (
	"testing"

	day02 "github.com/cpprian/AoC-2025/day02/lib"
)

func TestSolvePart1(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	data, err := day02.ParseIntoRange(input)
	if err != nil {
		t.Fail()
	}

	got := day02.SolvePart1(data)
	want := 1227775554
	if got != want {
		t.Errorf("got %d want %d given", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	data, err := day02.ParseIntoRange(input)
	if err != nil {
		t.Fail()
	}

	got := day02.SolvePart2(data)
	want := 4174379265
	if got != want {
		t.Errorf("got %d want %d given", got, want)
	}
}