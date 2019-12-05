package main

import (
	"fmt"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	util.Start()
	min, max := 246515, 739105

	count := 0
	for i := min; i <= max; i++ {
		hasDouble, isIncreasing := part1(i)
		if hasDouble && isIncreasing {
			count++
		}
	}

	fmt.Println("Part 1", count)
	util.Elapsed()

	util.Start()
	count = 0
	for i := min; i <= max; i++ {
		hasDouble, isIncreasing := part2(i)
		if hasDouble && isIncreasing {
			count++
		}
	}

	fmt.Println("Part 2", count)
	util.Elapsed()
}

func part1(number int) (bool, bool) {
	s := strconv.Itoa(number)
	var hasDouble, isIncreasing bool = false, true
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false, false
		}

		if s[i] == s[i-1] {
			hasDouble = true
		}
	}

	return hasDouble, isIncreasing
}

func part2(number int) (bool, bool) {
	s := strconv.Itoa(number)
	digits := make(map[byte]int)
	var hasDouble, isIncreasing bool = false, true
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false, false
		}
	}

	for i := 0; i < len(s); i++ {
		digits[s[i]] = digits[s[i]] + 1
	}

	for _, v := range digits {
		if v == 2 {
			hasDouble = true
		}
	}

	return hasDouble, isIncreasing
}
