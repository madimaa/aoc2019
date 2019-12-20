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
	intcode []int
	input   []int
	index   int
}

//CreateComputer - creats an intcode computer
func CreateComputer(intcode []int) *Computer {
	//making a copy of the slice will prevent modifying the `background array`
	intcodeCopy := make([]int, len(intcode))
	copy(intcodeCopy, intcode)
	return &Computer{intcode: intcodeCopy, input: make([]int, 0), index: 0}
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
			noun := getValue(computer.intcode, i+1, getDigit(num, 2))
			verb := getValue(computer.intcode, i+2, getDigit(num, 3))
			result := 0
			if opCode == 1 {
				result = noun + verb
			} else {
				result = noun * verb
			}

			putValue(computer.intcode, i+3, getDigit(num, 4), result)
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

			putValue(computer.intcode, i+1, getDigit(num, 2), number)
			computer.index += 2
		} else if opCode == 4 {
			output = getValue(computer.intcode, i+1, getDigit(num, 2))
			computer.index += 2
			break
		} else if opCode == 5 || opCode == 6 {
			firstParam := getValue(computer.intcode, i+1, getDigit(num, 2))
			if opCode == 5 && firstParam > 0 || opCode == 6 && firstParam == 0 {
				computer.index = getValue(computer.intcode, i+2, getDigit(num, 3))
			} else {
				computer.index += 3
			}
		} else if opCode == 7 || opCode == 8 {
			firstParam := getValue(computer.intcode, i+1, getDigit(num, 2))
			secondParam := getValue(computer.intcode, i+2, getDigit(num, 3))
			if opCode == 7 && firstParam < secondParam || opCode == 8 && firstParam == secondParam {
				putValue(computer.intcode, i+3, getDigit(num, 4), 1)
			} else {
				putValue(computer.intcode, i+3, getDigit(num, 4), 0)
			}

			computer.index += 4
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
