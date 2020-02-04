package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/intcode"
)

type point struct {
	x, y int
}

type robot struct {
	pos point
	dir int
}

const (
	north    = 0
	east     = 1
	south    = 2
	west     = 3
	space    = 46
	scaffold = 35
	newLine  = 10
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("17.txt")
	content := strings.Split(result[0], ",")
	intcodeArr := make([]int, 10000)
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	vacuumRobot, scaffolds := part1(intcodeArr)
	util.Elapsed()
	fmt.Println("Part 2")
	util.Start()
	part2(intcodeArr, scaffolds, vacuumRobot)
	util.Elapsed()
}

func part1(intcodeArr []int) (*robot, map[point]int) {
	computer := intcode.CreateComputer(intcodeArr)
	var output int
	halt := false

	journeyMap := make(map[point]int)
	position := point{0, 0}

	for !halt {
		_, output, halt = computer.Computer()
		journeyMap[position] = output
		if output == newLine {
			position.y++
			position.x = 0
			//fmt.Println()
		} else {
			position.x++
			//fmt.Print(string(output))
		}
	}

	vacuumRobot := &robot{point{0, 0}, north}
	scaffolds := make(map[point]int)
	alignmentSum := 0
	for k, v := range journeyMap {
		if v == scaffold {
			scaffolds[k] = v
			if checkIntersection(k, journeyMap) {
				alignmentSum += k.x * k.y
			}
		} else {
			value := string(v)
			switch value {
			case "^":
				vacuumRobot = &robot{k, north}
				break
			case ">":
				vacuumRobot = &robot{k, east}
				break
			case "v":
				vacuumRobot = &robot{k, south}
				break
			case "<":
				vacuumRobot = &robot{k, west}
				break
			}
		}
	}

	fmt.Println(alignmentSum)
	return vacuumRobot, scaffolds
}

func part2(intcodeArr []int, scaffolds map[point]int, vacuumRobot *robot) {
	path := calculatePath(scaffolds, vacuumRobot)
	mergedPath := strings.Join(path, "")
	a, b, c := breakIntoSeparateParts(path, mergedPath)

	command := strings.ReplaceAll(mergedPath, a, "A,")
	command = strings.ReplaceAll(command, b, "B,")
	command = strings.ReplaceAll(command, c, "C,")
	command = strings.Trim(command, ",")
	fmt.Println("Command:", command)

	a = splitSubcommand(a)
	b = splitSubcommand(b)
	c = splitSubcommand(c)
	fmt.Println("Subcommands:", a, b, c)

	intcodeArr[0] = 2
	computer := intcode.CreateComputer(intcodeArr)
	addIntcodeInput(computer, command)
	addIntcodeInput(computer, a)
	addIntcodeInput(computer, b)
	addIntcodeInput(computer, c)

	//I can not explain this.
	computer.AddInput(10)
	computer.AddInput(10)

	var output int
	halt := false

	largestOutput := 0
	for !halt {
		_, output, halt = computer.Computer()
		if output > largestOutput {
			largestOutput = output
		}
	}

	fmt.Println(largestOutput)
}

func checkIntersection(position point, journeyMap map[point]int) bool {
	north := journeyMap[point{position.x, position.y - 1}] == scaffold
	south := journeyMap[point{position.x, position.y + 1}] == scaffold
	west := journeyMap[point{position.x - 1, position.y}] == scaffold
	east := journeyMap[point{position.x + 1, position.y}] == scaffold

	return north && south && west && east
}

func calculatePath(scaffolds map[point]int, vacuumRobot *robot) []string {
	path := make([]string, 0)
	counter := 0

	for {
		if isScaffold(scaffolds, vacuumRobot) {
			counter++
			switch vacuumRobot.dir {
			case north:
				vacuumRobot.pos.y--
				break
			case south:
				vacuumRobot.pos.y++
				break
			case west:
				vacuumRobot.pos.x--
				break
			case east:
				vacuumRobot.pos.x++
				break
			}
		} else {
			canTurn, direction := turn(scaffolds, vacuumRobot)

			if canTurn {
				if counter == 0 {
					path = append(path, direction)
				} else {
					path = append(path, strconv.Itoa(counter))
					path = append(path, direction)
					counter = 0
				}
			} else {
				path = append(path, strconv.Itoa(counter))
				break
			}
		}
	}

	return path
}

func isScaffold(scaffolds map[point]int, vacuumRobot *robot) bool {
	ok := false
	switch vacuumRobot.dir {
	case north:
		_, ok = scaffolds[point{vacuumRobot.pos.x, vacuumRobot.pos.y - 1}]
		break
	case south:
		_, ok = scaffolds[point{vacuumRobot.pos.x, vacuumRobot.pos.y + 1}]
		break
	case west:
		_, ok = scaffolds[point{vacuumRobot.pos.x - 1, vacuumRobot.pos.y}]
		break
	case east:
		_, ok = scaffolds[point{vacuumRobot.pos.x + 1, vacuumRobot.pos.y}]
		break
	}

	return ok
}

func turn(scaffolds map[point]int, vacuumRobot *robot) (bool, string) {
	right := (vacuumRobot.dir + 1) % 4
	left := (vacuumRobot.dir + 3) % 4
	if isScaffold(scaffolds, &robot{vacuumRobot.pos, right}) {
		vacuumRobot.dir = right
		return true, "R"
	} else if isScaffold(scaffolds, &robot{vacuumRobot.pos, left}) {
		vacuumRobot.dir = left
		return true, "L"
	}

	return false, ""
}

func breakIntoSeparateParts(fullPath []string, path string) (string, string, string) {
	a := strings.Join(fullPath[0:2], "")
	var b, c string

	for i := 2; i < len(fullPath); i += 2 {
		parts := strings.Split(path, a)

		partB := parts[0]
		partC := parts[0]
		for j := 0; j < len(parts); j++ {
			actual := parts[j]
			if actual == "" {
				continue
			} else if partB == "" || len(partB) > len(actual) {
				partB = actual
			}
		}

		for j := 0; j < len(parts); j++ {
			actual := parts[j]
			if actual != partB {
				actualParts := strings.Split(actual, partB)
				partC = ""
				match := true
				for _, v := range actualParts {
					if v != "" && partC == "" {
						partC = v
					}

					if v != partC {
						match = false
					}
				}

				if match && partC != "" {
					break
				}
			}
		}

		partsFound := true
		for _, v := range parts {
			hasMatch := false
			str := strings.ReplaceAll(v, partB, "")
			str = strings.ReplaceAll(str, partC, "")
			if v == partB || v == partC || len(str) == 0 {
				hasMatch = true
			}

			if !hasMatch && v != "" {
				partsFound = false
				break
			}
		}

		if partsFound {
			b = partB
			c = partC
			break
		} else {
			a += strings.Join(fullPath[i:i+2], "")
		}
	}

	return a, b, c
}

func splitSubcommand(subcommand string) string {
	result := ""
	for i := 0; i < len(subcommand); i++ {
		char := string(subcommand[i])
		if char == "R" || char == "L" {
			result += ","

		}

		result += char
	}

	return strings.Trim(result, ",")
}

func addIntcodeInput(computer *intcode.Computer, command string) {
	for i := 0; i < len(command); i++ {
		computer.AddInput(int(command[i]))

		if command[i] == 'L' || command[i] == 'R' {
			computer.AddInput(',')
		}
	}

	computer.AddInput(10)
}
