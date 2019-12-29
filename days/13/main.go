package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/intcode"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("13.txt")
	content := strings.Split(result[0], ",")
	intcodeArr := make([]int, 10000)
	for i, s := range content {
		number, err := strconv.Atoi(s)
		util.LogOnError(err)
		intcodeArr[i] = number
	}

	part1(intcodeArr)
	part2(intcodeArr)

	util.Elapsed()
}

func part1(intcodeArr []int) {
	computer := intcode.CreateComputer(intcodeArr)
	var output int
	halt := false
	counter := 0

	pixels := make(map[string]int)

	var x, y int
	for !halt {
		_, output, halt = computer.Computer()
		switch counter % 3 {
		case 0:
			x = output
		case 1:
			y = output
		case 2:
			key := fmt.Sprintf("%d %d", x, y)
			pixels[key] = output
		}

		counter++
	}

	blockCounter := 0
	for _, v := range pixels {
		if v == 2 {
			blockCounter++
		}
	}

	printImage(pixels)
	fmt.Println(blockCounter)
}

type coordinates struct {
	x, y int
}

func part2(intcodeArr []int) {
	intcodeArr[0] = 2
	computer := intcode.CreateComputer(intcodeArr)
	var output int
	halt := false
	counter := 0
	points := 0

	actual := coordinates{0, 0}
	ball := coordinates{0, 0}
	paddle := coordinates{0, 0}
	for !halt {
		_, output, halt = computer.Computer()
		switch counter % 3 {
		case 0:
			actual.x = output
		case 1:
			actual.y = output
		case 2:
			switch output {
			case 0: //empty tile
			case 1: //wall
			case 2: //block
			case 3: //horizontal paddle
				paddle.x = actual.x
				paddle.y = actual.y
			case 4: //ball
				ball.x = actual.x
				ball.y = actual.y
				if ball.x == paddle.x {
					computer.AddInput(0)
				} else if ball.x > paddle.x {
					computer.AddInput(1)
				} else {
					computer.AddInput(-1)
				}
			default:
				points = output
			}
		}

		counter++
	}

	fmt.Println(points)
}

func printImage(pixels map[string]int) {
	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{37, 22}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for k, v := range pixels {
		var x, y int
		fmt.Sscanf(k, "%d %d", &x, &y)
		switch v {
		case 0: //empty tile
			img.Set(x, y, color.Black)
		case 1: //wall
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		case 2: //block
			img.Set(x, y, color.RGBA{255, 0, 255, 255})
		case 3: //horizontal paddle
			img.Set(x, y, color.RGBA{255, 255, 0, 255})
		case 4: //ball
			img.Set(x, y, color.RGBA{122, 122, 0, 255})
		}
	}

	file, err := os.Create("13.png")
	util.PanicOnError(err)
	png.Encode(file, img)
}
