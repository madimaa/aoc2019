package main

import (
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")

	result := util.OpenFile("01.txt")
	var sum int = 0
	for _, s := range result {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		sum += calculate(number)
	}

	fmt.Println(sum)

	fmt.Println("Part 2")
	sum = 0
	for _, s := range result {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		actual := calculate(number)
		sum += actual
		for actual > 0 {
			if actual = calculate(actual); actual > 0 {
				sum += actual
			}
		}
	}

	fmt.Println(sum)

	os.Exit(0)
}

func calculate(number int) int {
	return number/3 - 2
}
