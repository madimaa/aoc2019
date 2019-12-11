package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/intcode"
)

func main() {
	fmt.Println("Part 1")
	fmt.Println("Input = 1")
	util.Start()
	result := util.OpenFile("05.txt")
	content := strings.Split(result[0], ",")
	// result := "1002,4,3,4,33"
	// content := strings.Split(result, ",")
	intcodeArr := make([]int, len(content))
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	computer := intcode.CreateComputer(intcodeArr, true, 0)

	fmt.Println(computer.Computer())
	util.Elapsed()

	fmt.Println("Part 2")
	fmt.Println("Input = 5")
	util.Start()
	fmt.Println(computer.Computer())
	util.Elapsed()
}

func computer(input []int) int {
	//making a copy of the slice will prevent modifying the `background array`
	intcode := make([]int, len(input))
	copy(intcode, input)

	for i := 0; i < len(intcode); {
		num := intcode[i]
		opCode := getOpCode(num)
		//fmt.Println(opCode)
		if opCode == 99 {
			break
		}

		if opCode == 1 || opCode == 2 {
			//fmt.Println(getDigit(num, 2))
			//fmt.Println(getDigit(num, 3))
			noun := getValue(intcode, i+1, getDigit(num, 2))
			verb := getValue(intcode, i+2, getDigit(num, 3))
			result := 0
			if opCode == 1 {
				result = noun + verb
			} else {
				result = noun * verb
			}

			putValue(intcode, i+3, getDigit(num, 4), result)
			i += 4
		} else if opCode == 3 {
			reader := bufio.NewReader(os.Stdin)
			text, _, err := reader.ReadLine()
			util.LogOnError(err)
			num, err := strconv.Atoi(string(text))
			util.PanicOnError(err)
			putValue(intcode, i+1, getDigit(num, 2), num)
			i += 2
		} else if opCode == 4 {
			fmt.Println(getValue(intcode, i+1, getDigit(num, 2)))
			i += 2
		} else if opCode == 5 || opCode == 6 {
			firstParam := getValue(intcode, i+1, getDigit(num, 2))
			if opCode == 5 && firstParam > 0 || opCode == 6 && firstParam == 0 {
				i = getValue(intcode, i+2, getDigit(num, 3))
			} else {
				i += 3
			}
		} else if opCode == 7 || opCode == 8 {
			firstParam := getValue(intcode, i+1, getDigit(num, 2))
			secondParam := getValue(intcode, i+2, getDigit(num, 3))
			if opCode == 7 && firstParam < secondParam || opCode == 8 && firstParam == secondParam {
				putValue(intcode, i+3, getDigit(num, 4), 1)
			} else {
				putValue(intcode, i+3, getDigit(num, 4), 0)
			}

			i += 4
		}
	}

	return intcode[0]
}

func getOpCode(input int) int {
	if input > 99 {
		return getDigits(input, 0, 2)
	}

	return input
}

func getDigit(input, from int) int {
	//fmt.Println(input, from)
	return getDigits(input, from, from+1)
}

func getDigits(input, from, to int) int {
	s := strconv.Itoa(input)
	length := len(s)
	if (length == from || length < from) && length < to {
		return 0
	} else if length >= from && length < to {
		num, err := strconv.Atoi(s[from:length])
		util.PanicOnError(err)
		return num
	} else {
		num, err := strconv.Atoi(s[length-to : length-from])
		util.PanicOnError(err)
		return num
	}
}

func getValue(intcode []int, position, mode int) int {
	switch mode {
	case 0:
		return intcode[intcode[position]]
	case 1:
		return intcode[position]
	default:
		panic("Something went wrong")
	}
}

func putValue(intcode []int, position, mode, value int) {
	switch mode {
	case 0:
		intcode[intcode[position]] = value
	case 1:
		intcode[position] = value
	default:
		panic("Something went wrong")
	}
}
