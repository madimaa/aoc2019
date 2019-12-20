package main

import (
	"fmt"
	"strconv"
	"strings"

	prmt "github.com/gitchander/permutation"
	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/intcode"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("07.txt")
	content := strings.Split(result[0], ",")
	// result := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	// content := strings.Split(result, ",")
	intcodeArr := make([]int, len(content))
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	maxOutput := 0
	elements := []int{0, 1, 2, 3, 4}
	permutations := prmt.New(prmt.IntSlice(elements))
	for permutations.Next() {
		a, b, c, d, e := elements[0], elements[1], elements[2], elements[3], elements[4]
		aOut := compute(intcodeArr, a, 0)
		bOut := compute(intcodeArr, b, aOut)
		cOut := compute(intcodeArr, c, bOut)
		dOut := compute(intcodeArr, d, cOut)
		out := compute(intcodeArr, e, dOut)

		if maxOutput < out {
			maxOutput = out
		}
	}

	fmt.Println(maxOutput)
	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	maxOutput = 0
	elements = []int{5, 6, 7, 8, 9}
	permutations = prmt.New(prmt.IntSlice(elements))
	for permutations.Next() {
		a := intcode.CreateComputer(intcodeArr)
		a.AddInput(elements[0])
		b := intcode.CreateComputer(intcodeArr)
		b.AddInput(elements[1])
		c := intcode.CreateComputer(intcodeArr)
		c.AddInput(elements[2])
		d := intcode.CreateComputer(intcodeArr)
		d.AddInput(elements[3])
		e := intcode.CreateComputer(intcodeArr)
		e.AddInput(elements[4])
		halt := false
		nextInput := 0
		amplifier := 0
		var computer *intcode.Computer
		for !halt {
			switch amplifier % 5 {
			case 0:
				computer = a
			case 1:
				computer = b
			case 2:
				computer = c
			case 3:
				computer = d
			case 4:
				computer = e
			}
			nextInput, halt = computeWithPersistedData(computer, nextInput)
			amplifier++
			if amplifier%5 == 0 {
				if maxOutput < nextInput {
					maxOutput = nextInput
				}
			}
		}
	}

	fmt.Println(maxOutput)
	util.Elapsed()
}

func compute(intcodeArr []int, phaseSettings, input int) int {
	computer := intcode.CreateComputer(intcodeArr)
	computer.AddInput(phaseSettings)
	computer.AddInput(input)
	_, output, _ := computer.Computer()
	return output
}

func computeWithPersistedData(computer *intcode.Computer, input int) (int, bool) {
	computer.AddInput(input)
	_, output, halt := computer.Computer()
	return output, halt
}
