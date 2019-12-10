package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

//ComputerWithInput - add noun and verb
func ComputerWithInput(input []int, noun, verb int) int {
	//making a copy of the slice will prevent modifying the `background array`
	intcode := make([]int, len(input))
	copy(intcode, input)
	intcode[1] = noun
	intcode[2] = verb
	return Computer(intcode)
}

//Computer - intcode computer
func Computer(input []int) int {
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
