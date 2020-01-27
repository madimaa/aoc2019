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

type point struct {
	x, y int
}

const (
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

	part1(intcodeArr)
	util.Elapsed()
	fmt.Println("Part 2")
	util.Start()
	part2(intcodeArr)
	util.Elapsed()
}

func part1(intcodeArr []int) {
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

	alignmentSum := 0
	for k, v := range journeyMap {
		if v == scaffold && checkIntersection(k, journeyMap) {
			alignmentSum += k.x * k.y
		}
	}

	fmt.Println(alignmentSum)
}

func part2(intcodeArr []int) {

}

func checkIntersection(position point, journeyMap map[point]int) bool {
	north := journeyMap[point{position.x, position.y - 1}] == scaffold
	south := journeyMap[point{position.x, position.y + 1}] == scaffold
	west := journeyMap[point{position.x - 1, position.y}] == scaffold
	east := journeyMap[point{position.x + 1, position.y}] == scaffold

	return north && south && west && east
}

func printImage(pixels map[string]string) {
	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{100, 100}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for k, v := range pixels {
		var x, y int
		fmt.Sscanf(k, "%d %d", &x, &y)

		x += 50
		y += 50

		if x == 50 && y == 50 {
			img.Set(x, y, color.RGBA{0, 255, 0, 255})
			continue
		}

		switch v {
		// case wall:
		// 	img.Set(x, y, color.Black)
		// case path, deadEnd:
		// 	img.Set(x, y, color.White)
		// case oxygen:
		// 	img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	file, err := os.Create("15.png")
	util.PanicOnError(err)
	png.Encode(file, img)
}
