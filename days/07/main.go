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
	// result := "1002,4,3,4,33"
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
		computer := intcode.CreateComputer(intcodeArr)
		computer.AddInput(a)
		computer.AddInput(0)
		_, aOut := computer.Computer()

		computer = intcode.CreateComputer(intcodeArr)
		computer.AddInput(b)
		computer.AddInput(aOut)
		_, bOut := computer.Computer()

		computer = intcode.CreateComputer(intcodeArr)
		computer.AddInput(c)
		computer.AddInput(bOut)
		_, cOut := computer.Computer()

		computer = intcode.CreateComputer(intcodeArr)
		computer.AddInput(d)
		computer.AddInput(cOut)
		_, dOut := computer.Computer()

		computer = intcode.CreateComputer(intcodeArr)
		computer.AddInput(e)
		computer.AddInput(dOut)
		_, out := computer.Computer()

		if maxOutput < out {
			maxOutput = out
		}
	}

	fmt.Println(maxOutput)
	util.Elapsed()
}

func splitQuinaryFromNumber(number int64) (int64, int64, int64, int64, int64) {
	base := 5
	quinaryString := strconv.FormatInt(number, base)
	//fmt.Println(quinaryString)
	length := len(quinaryString)

	var a, b, c, d, e int64 = 0, 0, 0, 0, 0
	var err error

	if length >= 1 {
		e, err = strconv.ParseInt(quinaryString[length-1:length], base, 64)
		util.PanicOnError(err)
	}

	if length >= 2 {
		num := quinaryString[length-2 : length-1]
		d, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 3 {
		num := quinaryString[length-3 : length-2]
		c, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 4 {
		num := quinaryString[length-4 : length-3]
		b, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	if length >= 5 {
		num := quinaryString[length-5 : length-4]
		a, err = strconv.ParseInt(num, base, 64)
		util.PanicOnError(err)
	}

	return a, b, c, d, e
}
