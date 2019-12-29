package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2019/lib/array2d"

	"github.com/madimaa/aoc2019/lib/intcode"

	util "github.com/madimaa/aoc2019/lib"
)

type vector struct {
	x, y, direction int
}

func createVector(x, y, direction int) *vector {
	return &vector{x: x, y: y, direction: direction}
}

var array *array2d.Array2D
var robot *vector
var painted map[string]bool

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("11.txt")
	content := strings.Split(result[0], ",")
	intcodeArr := make([]int, 10000)
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	painted = make(map[string]bool)
	// painted = make([]int, 0)

	size := 140
	array = array2d.Create(size, size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			array.Put(x, y, ".")
		}
	}

	//directions: 0 up, 1 right, 2 down, 3 left
	robot = createVector(size/2, size/2, 0)

	/*
		All of the panels are currently black. (provide 0 if the robot is over a black panel or 1 if the robot is over a white panel)
	*/
	runComputer(0, intcodeArr) //first input. every panel is black

	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	//backpaint
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			array.Put(x, y, ".")
		}
	}

	array.Put(robot.x, robot.y, "#") //first input. every panel is white

	runComputer(1, intcodeArr) //first input. every panel is white

	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			str := array.Get(x, y)
			if str == "." {
				img.Set(x, y, color.Black)
			} else if str == "#" {
				img.Set(x, y, color.White)
			}
		}
	}

	file, err := os.Create("11.png")
	util.PanicOnError(err)
	png.Encode(file, img)

	fmt.Println(len(painted))
	util.Elapsed()
}

func runComputer(input int, intcodeArr []int) {
	computer := intcode.CreateComputer(intcodeArr)
	computer.AddInput(input)
	var output int
	halt := false
	counter := 0
	for !halt {
		_, output, halt = computer.Computer()
		if counter%2 == 0 {
			if output == 0 {
				array.Put(robot.x, robot.y, ".")
			} else {
				array.Put(robot.x, robot.y, "#")
			}

			painted[fmt.Sprintf("%d%d", robot.x, robot.y)] = true
		} else {
			turnRobot(output)
			moveRobot()

			if array.Get(robot.x, robot.y) == "." {
				computer.AddInput(0)
			} else {
				computer.AddInput(1)
			}
		}

		counter++
	}
}

//directions: 0 up, 1 right, 2 down, 3 left
func turnRobot(turn int) {
	if turn == 1 {
		robot.direction++
	} else {
		robot.direction--
	}

	robot.direction += 4
	robot.direction %= 4
}

func moveRobot() {
	switch robot.direction {
	case 0:
		robot.y--
	case 1:
		robot.x++
	case 2:
		robot.y++
	case 3:
		robot.x--
	}
}
