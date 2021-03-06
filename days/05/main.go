package main

import (
	"fmt"
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

	computer := intcode.CreateComputer(intcodeArr)
	computer.AddInput(1)
	var output int
	_, output, _ = computer.Computer()
	//fix after output break in day 07 part 2
	for output == 0 {
		_, output, _ = computer.Computer()
	}
	fmt.Println(output)
	util.Elapsed()

	fmt.Println("Part 2")
	fmt.Println("Input = 5")
	util.Start()
	computer = intcode.CreateComputer(intcodeArr)
	computer.AddInput(5)
	fmt.Println(computer.Computer())
	util.Elapsed()
}
