package day02

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var rangeRe = regexp.MustCompile(`[1-9]\d*-[1-9]\d*`)

type Range struct {
	left, right int
}

type Ranges []Range

func pow10(k int) int64 {
	res := int64(1)
	for i := 0; i < k; i++ {
		res *= 10
	}
	return res
}

// too low 19605386012
func (r Range) solveRange1() int {
	var sum int = 0

	left := r.left
	right := r.right

	maxLen := int(math.Log10(float64(right))) + 1

	for k := 1; k * 2 <= maxLen; k++ {
		p := pow10(k)            
		d := p + 1             
		minN := (left + int(d) - 1) / int(d) 
		maxN := right / int(d)            

		lowN := pow10(k-1)
		highN := p - 1

		if minN < int(lowN) {
			minN = int(lowN)
		}
		if maxN > int(highN) {
			maxN = int(highN)
		}
		if minN > maxN {
			continue
		}

		count := maxN - minN + 1
		sumN := (minN + maxN) * count / 2
		sum += int(d) * sumN
	}

	return sum
}

func (r Range) solveRange2() int {
	var sum int = 0

	left := r.left
	right := r.right

	maxLen := int(math.Log10(float64(right))) + 1

	for k := 1; k * 2 <= maxLen; k++ {
		p := pow10(k)            
		d := p + 1             
		minN := (left + int(d) - 1) / int(d) 
		maxN := right / int(d)            

		lowN := pow10(k-1)
		highN := p - 1

		if minN < int(lowN) {
			minN = int(lowN)
		}
		if maxN > int(highN) {
			maxN = int(highN)
		}
		if minN > maxN {
			continue
		}

		count := maxN - minN + 1
		sumN := (minN + maxN) * count / 2
		sum += int(d) * sumN
	}

	return sum
}

func SolvePart1(rgs Ranges) int {
	var solution int
	for _, r := range rgs {
		solution += r.solveRange1()
	}
	return solution
}

func SolvePart2(rgs Ranges) int {
	var solution int
	for _, r := range rgs {
		solution += r.solveRange2()
	}
	return solution
}

func parseIntoInt64(s string) (int, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("Cannot parse into int64: %v\n", err)
		return 0, err
	}
	return int(num), nil
}

func parseIntoStruct(s string) (Range, error) {
	newRange := strings.Split(s, "-")
	left, err := parseIntoInt64(newRange[0])
	if err != nil {
		return Range{}, err
	}	

	right, err := parseIntoInt64(newRange[1])
	if err != nil {
		return Range{}, err
	}

	if left > right {
		return Range{}, fmt.Errorf("left range is larger than right")
	}
	return Range{left, right}, nil
}

func ParseIntoRange(input string) (Ranges, error) {
	var ranges = Ranges{}

	matches := rangeRe.FindAllString(input, -1)
	for _, match := range matches {
		newRange, err := parseIntoStruct(match)
		if err != nil {
			return Ranges{}, err
		}
		ranges = append(ranges, newRange)
	}
	return ranges, nil
}