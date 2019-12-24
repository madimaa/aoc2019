package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
)

//Computer - intcode computer
type Computer struct {
	intcode      []int
	input        []int
	index        int
	relativeBase int
}

//CreateComputer - creats an intcode computer
func CreateComputer(intcode []int) *Computer {
	//making a copy of the slice will prevent modifying the `background array`
	intcodeCopy := make([]int, len(intcode))
	copy(intcodeCopy, intcode)
	return &Computer{intcode: intcodeCopy, input: make([]int, 0), index: 0, relativeBase: 0}
}

//Print - prints the intcode array
func (computer *Computer) Print() {
	fmt.Println(computer.intcode)
}

//AddInput - add input to fifo
func (computer *Computer) AddInput(input int) {
	computer.input = append(computer.input, input)
}

//ComputerWithInput - add noun and verb
func (computer *Computer) ComputerWithInput(noun, verb int) (int, int, bool) {
	computer.intcode[1] = noun
	computer.intcode[2] = verb
	return computer.Computer()
}

//Computer - intcode computer
func (computer *Computer) Computer() (int, int, bool) {
	halt := false
	output := 0
	for computer.index < len(computer.intcode) {
		i := computer.index
		num := computer.intcode[i]
		opCode := getOpCode(num)
		//fmt.Println(opCode)
		if opCode == 99 {
			halt = true
			break
		}

		if opCode == 1 || opCode == 2 {
			noun := getValue(computer.intcode, i+1, getDigit(num, 2), computer.relativeBase)
			verb := getValue(computer.intcode, i+2, getDigit(num, 3), computer.relativeBase)
			result := 0
			if opCode == 1 {
				result = noun + verb
			} else {
				result = noun * verb
			}

			putValue(computer.intcode, i+3, getDigit(num, 4), result, computer.relativeBase)
			computer.index += 4
		} else if opCode == 3 {
			var number int
			if len(computer.input) == 0 {
				fmt.Println(output)
				reader := bufio.NewReader(os.Stdin)
				text, _, err := reader.ReadLine()
				util.LogOnError(err)
				number, err = strconv.Atoi(string(text))
				util.PanicOnError(err)
			} else {
				number = computer.input[0]
				computer.input = computer.input[1:]
			}

			putValue(computer.intcode, i+1, getDigit(num, 2), number, computer.relativeBase)
			computer.index += 2
		} else if opCode == 4 {
			output = getValue(computer.intcode, i+1, getDigit(num, 2), computer.relativeBase)
			computer.index += 2
			break
		} else if opCode == 5 || opCode == 6 {
			firstParam := getValue(computer.intcode, i+1, getDigit(num, 2), computer.relativeBase)
			if opCode == 5 && firstParam > 0 || opCode == 6 && firstParam == 0 {
				computer.index = getValue(computer.intcode, i+2, getDigit(num, 3), computer.relativeBase)
			} else {
				computer.index += 3
			}
		} else if opCode == 7 || opCode == 8 {
			firstParam := getValue(computer.intcode, i+1, getDigit(num, 2), computer.relativeBase)
			secondParam := getValue(computer.intcode, i+2, getDigit(num, 3), computer.relativeBase)
			if opCode == 7 && firstParam < secondParam || opCode == 8 && firstParam == secondParam {
				putValue(computer.intcode, i+3, getDigit(num, 4), 1, computer.relativeBase)
			} else {
				putValue(computer.intcode, i+3, getDigit(num, 4), 0, computer.relativeBase)
			}

			computer.index += 4
		} else if opCode == 9 {
			valuePosition := 0
			switch getDigit(num, 2) {
			case 0: //position mode
				valuePosition = computer.intcode[i+1]
			case 1: //immediate mode
				valuePosition = i + 1
			case 2: //relative mode
				valuePosition = computer.relativeBase + computer.intcode[i+1]
			default:
				panic("Something went wrong")
			}

			computer.relativeBase += computer.intcode[valuePosition]
			computer.index += 2
		}
	}

	return computer.intcode[0], output, halt
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

func getValue(intcode []int, position, mode, relativeBase int) int {
	switch mode {
	case 0: //position mode
		return intcode[intcode[position]]
	case 1: //immediate mode
		return intcode[position]
	case 2: //relative mode
		return intcode[relativeBase+intcode[position]]
	default:
		panic("Something went wrong")
	}
}

func putValue(intcode []int, position, mode, value, relativeBase int) {
	switch mode {
	case 0: //position mode
		intcode[intcode[position]] = value
	case 1: //immediate mode
		intcode[position] = value
	case 2: //relative mode
		intcode[relativeBase+intcode[position]] = value
	default:
		panic("Something went wrong")
	}
}
