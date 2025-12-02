package main

import (
	"fmt"
	"os"

	day02 "github.com/cpprian/AoC-2025/day02/lib"
)

func main() {
	data, err := os.ReadFile("../inputs/day02")
	if err != nil {
		fmt.Printf("Cannot open file: %v\n", err)
		os.Exit(1)
	}

	input, err := day02.ParseIntoRange(string(data))
	if err != nil {
		fmt.Printf("Cannot solve challenges due to: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Solution part 1: %d\n", day02.SolvePart1(input))
	fmt.Printf("Solution part 2: %d\n", day02.SolvePart2(input))
}
