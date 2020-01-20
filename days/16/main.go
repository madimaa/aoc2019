package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("sample3.txt")
	input := result[0]
	pattern := [4]int{0, 1, 0, -1}

	patterns := calculatePatterns(pattern, len(input))

	phases := 100

	output := input
	for i := 1; i <= phases; i++ {
		newInput := ""
		for _, actualPattern := range patterns {
			newInput += fft(i, output, actualPattern)
		}

		output = newInput
	}

	fmt.Println(output[0:8])
	util.Elapsed()
}

func calculatePatterns(pattern [4]int, length int) []*list.List {
	patterns := make([]*list.List, 0)

	for i := 1; i <= length; i++ {
		patterns = append(patterns, calculateNextPattern(pattern, length, i))
	}

	return patterns
}

func fft(phase int, input string, pattern *list.List) string {
	output := ""

	j := 0

	result := 0
	for e := pattern.Front(); e != nil; e = e.Next() {
		num, err := strconv.Atoi(string(input[j]))
		util.PanicOnError(err)

		result += num * e.Value.(int)

		j++
	}

	val := math.Abs(float64(result % 10))
	output += strconv.Itoa(int(val))

	return output
}

func calculateNextPattern(pattern [4]int, length, serialNumber int) *list.List {
	actualPattern := list.New()
	counter := 0
	for actualPattern.Len() <= length {
		for n := 0; n < serialNumber; n++ {
			actualPattern.PushBack(pattern[counter%4])
		}

		counter++
	}

	actualPattern.Remove(actualPattern.Front())

	for actualPattern.Len() > length {
		actualPattern.Remove(actualPattern.Back())
	}

	return actualPattern
}
