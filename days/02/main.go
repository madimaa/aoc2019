package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019"
)

func main() {
	fmt.Println("Part 1")
	util.Start()
	result := util.OpenFile("02.txt")
	content := strings.Split(result[0], ",")
	intcode := make([]int, len(content))
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcode[i] = number
	}

	fmt.Println(computer(intcode, 12, 2))
	util.Elapsed()

	fmt.Println("Part 2")
	partTwo(intcode)

	os.Exit(0)
}

func partTwo(intcode []int) {
	util.Start()
	goal := 19690720
	var noun, verb int = 0, 0

	zero := computer(intcode, noun, verb)
	nounSecond := computer(intcode, noun+1, verb)
	verbSecond := computer(intcode, noun, verb+1)

	nounIncrement := nounSecond - zero
	verbIncrement := verbSecond - zero

	if nounIncrement > verbIncrement {
		noun = goal / nounIncrement
		verb = (goal - zero - noun*nounIncrement) / verbIncrement
	} else {
		verb = goal / verbIncrement
		noun = (goal - zero - verb*verbIncrement) / nounIncrement
	}

	fmt.Println("Noun:", noun, "Verb:", verb)
	util.Elapsed()
}

func computer(input []int, noun, verb int) int {
	//making a copy of the slice will prevent modifying the `background array`
	intcode := make([]int, len(input))
	copy(intcode, input)

	intcode[1] = noun
	intcode[2] = verb
	for i := 0; i < len(intcode); i++ {
		num := intcode[i]
		if num == 99 {
			break
		}

		firstpos := intcode[i+1]
		secondpos := intcode[i+2]
		resultpos := intcode[i+3]
		switch num {
		case 1:
			intcode[resultpos] = intcode[firstpos] + intcode[secondpos]
		case 2:
			intcode[resultpos] = intcode[firstpos] * intcode[secondpos]
		}

		i += 3
	}

	return intcode[0]
}
