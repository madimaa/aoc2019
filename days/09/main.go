package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2019/lib/intcode"

	util "github.com/madimaa/aoc2019/lib"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	// result := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	// content := strings.Split(result, ",")
	result := util.OpenFile("09.txt")
	content := strings.Split(result[0], ",")
	intcodeArr := make([]int, 10000)
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	runComputer(1, intcodeArr)

	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	runComputer(2, intcodeArr)

	util.Elapsed()
}

func runComputer(input int, intcodeArr []int) {
	computer := intcode.CreateComputer(intcodeArr)
	computer.AddInput(input)
	var output, ret int
	halt := false
	for !halt {
		ret, output, halt = computer.Computer()
		fmt.Println(output, ret)
	}
}
