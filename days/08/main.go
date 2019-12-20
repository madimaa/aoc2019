package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2019/lib"
	"github.com/madimaa/aoc2019/lib/array2d"
)

func main() {
	fmt.Println("Part 1")
	util.Start()

	result := util.OpenFile("08.txt")
	content := result[0]
	width, height := 25, 6

	fmt.Println(part1(content, width, height))
	util.Elapsed()

	fmt.Println("Part 2")
	util.Start()

	part2(content, width, height)
	util.Elapsed()
}

func part1(content string, width, height int) int {
	row := 0
	layer := 0
	fewerZeros, fewerOnes, fewerTwos := 0, 0, 0
	actualZeros, actualOnes, actualTwos := 0, 0, 0
	for i := 0; i < len(content); i++ {
		if i%width == 0 {
			row++

			if row > height {
				row = 1
				layer++
				if fewerZeros == 0 || fewerZeros > actualZeros {
					fewerZeros = actualZeros
					fewerOnes = actualOnes
					fewerTwos = actualTwos
				}

				actualZeros, actualOnes, actualTwos = 0, 0, 0
			}
		}

		number, err := strconv.Atoi(string(content[i]))
		util.PanicOnError(err)
		switch number {
		case 0:
			actualZeros++
		case 1:
			actualOnes++
		case 2:
			actualTwos++
		}
	}

	return fewerOnes * fewerTwos
}

func part2(content string, width, height int) {
	row := -1
	layer := 0

	array := array2d.Create(width, height)

	for i := 0; i < len(content); i++ {
		//fmt.Println(i)
		x := i % width
		if x == 0 {
			row++

			if row == height {
				row = 0
				layer++
			}
		}

		number, err := strconv.Atoi(string(content[i]))
		util.PanicOnError(err)
		//fmt.Println(x)
		//fmt.Println(row)
		putPixel(array, x, row, number)
	}

	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			str := array.Get(x, y)
			if str == "0" {
				img.Set(x, y, color.Black)
			} else if str == "1" {
				img.Set(x, y, color.White)
			}

			fmt.Print(str)
		}

		fmt.Println()
	}

	file, err := os.Create("08.png")
	util.PanicOnError(err)
	png.Encode(file, img)
}

func putPixel(array *array2d.Array2D, x, y, number int) {
	if array.Get(x, y) == "" || array.Get(x, y) == "2" {
		array.Put(x, y, strconv.Itoa(number))
	}
}
