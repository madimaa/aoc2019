package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("16.txt")
	input := result[0]
	pattern := [4]int{0, 1, 0, -1}

	patterns := calculatePatterns(pattern, len(input))

	phases := 100

	output := input
	for phase := 1; phase <= phases; phase++ {
		newInput := ""
		for i := 1; i <= len(input); i++ {
			newInput += fft(phase, output, patterns[i])
		}

		output = newInput
	}

	fmt.Println(output[0:8])
	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	fmt.Println(part2(input, phases))

	util.Elapsed()
}

func calculatePatterns(pattern [4]int, length int) map[int][]int {
	patterns := make(map[int][]int)

	for i := 1; i <= length; i++ {
		patterns[i] = calculateNextPattern(pattern, length, i)
	}

	return patterns
}

func fft(phase int, input string, pattern []int) string {
	output := ""

	j := 0

	result := 0
	for _, value := range pattern {
		num, err := strconv.Atoi(string(input[j]))
		util.PanicOnError(err)

		result += num * value

		j++
	}

	val := math.Abs(float64(result % 10))
	output += strconv.Itoa(int(val))

	return output
}

func calculateNextPattern(pattern [4]int, length, serialNumber int) []int {
	actualPattern := make([]int, 0)
	counter := 0
	for len(actualPattern) <= length {
		for n := 0; n < serialNumber; n++ {
			actualPattern = append(actualPattern, pattern[counter%4])
		}

		counter++
	}

	actualPattern = actualPattern[1:] //remove first element

	for len(actualPattern) > length {
		actualPattern = actualPattern[:len(actualPattern)-1]
	}

	return actualPattern
}

func part2(input string, phases int) []int {
	part2Input := strings.Repeat(input, 10000)

	offset, err := strconv.Atoi(input[0:7])
	util.PanicOnError(err)

	newInput := make([]int, 0)

	for i := offset; i < len(part2Input); i++ {
		num, err := strconv.Atoi(string(part2Input[i]))
		util.PanicOnError(err)
		newInput = append(newInput, num)
	}

	for phase := 1; phase <= phases; phase++ {
		newInput = secondHalf(newInput)
	}

	return newInput[0:8]
}

func secondHalf(input []int) []int {
	output := make([]int, 0)
	sum := 0
	for i := len(input) - 1; i >= 0; i-- {
		num := input[i]
		sum += num
		output = append(output, sum%10)
	}

	for left, right := 0, len(output)-1; left < right; left, right = left+1, right-1 {
		output[left], output[right] = output[right], output[left]
	}

	return output
}
