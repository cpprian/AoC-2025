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

func (r Range) solveRange() int {
	if r.left == 0 {
		r.left = 1
	}
	el := strconv.Itoa(int(r.left))

	var counter int
	for {
		n, _ := strconv.Atoi(el)
		if len(el) % 2 == 1 {
			el += "0"
			n = int(math.Pow(10, float64(len(el) - 1)))
			if n > r.right {
				break
			}
		}

		p := int(math.Pow(10, float64(len(el)/2)))
		n_l := n / p
		n_r := n % p
		if n_r == 0 || len(strconv.Itoa(n_r)) < len(strconv.Itoa(n_l)) {
			n_r = n_l
			if n_l * p + n_r > r.right {
				break
			}
		}

		if n_l != n_r {
			n_r = n_l
			if n_l * p + n_r > r.right {
				
			}
			counter += 1
		}

		fmt.Println(r.left, r.right, n, n_l, n_r, counter)
		break
	}
	return counter
}

func SolvePart1(rgs Ranges) int {
	var solution int
	for _, r := range rgs {
		solution += r.solveRange()
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